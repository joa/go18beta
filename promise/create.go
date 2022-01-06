package promise

import (
	"sync/atomic"
	"unsafe"

	"github.com/joa/go18beta/attempt"
	"github.com/joa/go18beta/future"
	"github.com/joa/go18beta/option"
)

const (
	promiseInit    = 0 // promise is initialized
	promiseDone    = 1 // promise is done
	promiseWriting = 2 // promise is being written
)

func swapCallbacksForValue[T any](statePtr *int32, resPtr *unsafe.Pointer, res attempt.Attempt[T]) (*callback[T], bool) {
	for {
		switch oldState := atomic.LoadInt32(statePtr); oldState {
		case promiseInit:
			// current state is init, and we should transition to the done
			// state. since we're updating two atomics we enter a writing
			// state first.
			if !atomic.CompareAndSwapInt32(statePtr, oldState, promiseWriting) {
				// we lost the race and can't transition into writing state, retry
				continue
			} // else we won the race and will continue below
		case promiseDone:
			return nil, false // already done
		case promiseWriting:
			continue // someone else is updating, retry
		}

		// we were able to transition from init to writing
		// and are therefore able to update the resPtr

		callbacks := atomic.LoadPointer(resPtr)

		atomic.StorePointer(resPtr, unsafe.Pointer(&res))
		atomic.StoreInt32(statePtr, promiseDone)

		return (*callback[T])(callbacks), true
	}
}

func Create[T any]() Promise[T] {
	var state int32
	var result unsafe.Pointer

	atomic.StoreInt32(&state, promiseInit)
	atomic.StorePointer(&result, unsafe.Pointer(nilCallback))

	p := new(prom[T])

	p.doneFunc = func() bool { return atomic.LoadInt32(&state) == promiseDone }

	p.valueFunc = func() option.Option[attempt.Attempt[T]] {
		if atomic.LoadInt32(&state) == promiseDone {
			res := *((*attempt.Attempt[T])(atomic.LoadPointer(&result)))
			return option.Some(res)
		}

		return option.None[attempt.Attempt[T]]()
	}

	p.onCompleteFunc = func(f func(attempt.Attempt[T])) future.Future[T] {
		for {
			switch oldState := atomic.LoadInt32(&state); oldState {
			case promiseInit:
				// current state is init, and we should transition to the writing
				// state. this protects us from other writers while updating the
				// list of callbacks
				if !atomic.CompareAndSwapInt32(&state, oldState, promiseWriting) {
					// we lost the race and can't transition into writing state, retry
					continue
				} // else we won the race and will continue below
			case promiseDone:
				res := *((*attempt.Attempt[T])(atomic.LoadPointer(&result)))
				f(res)
				return p.Future()
			case promiseWriting:
				continue // someone else is updating, retry
			}

			newCallbacks := &callback[T]{
				f:     f,
				next:  (*callback[T])(atomic.LoadPointer(&result)),
				value: new(atomic.Value),
			}

			atomic.StorePointer(&result, unsafe.Pointer(newCallbacks))
			atomic.StoreInt32(&state, promiseInit)

			return p.Future()
		}
	}

	p.tryCompleteFunc = func(a attempt.Attempt[T]) bool {
		switch cbs, changed := swapCallbacksForValue(&state, &result, a); {
		case !changed:
			return false // already completed
		case cbs == nil:
			return true // successfully completed without listeners
		default:
			// successfully completed with listeners
			cb := reverseCallbackListAndRemoveNil[T](cbs)
			for cb != nil {
				next := cb.next
				cb.dispatch(a)
				cb = next
			}
			return true
		}
	}

	return p
}
