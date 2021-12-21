package promise

import (
	"fmt"
	"github.com/joa/go18beta/attempt"
	"github.com/joa/go18beta/future"
	"github.com/joa/go18beta/option"
	"sync/atomic"
	"unsafe"
)

func swapCallbacksForValue[T any](statePtr *unsafe.Pointer, donePtr *int32, newState attempt.Attempt[T]) *callback[T] {
	for i := 0; i < 1000; i++ {
		// TODO: this is not safe, needs a write-mark for the donePtr
		alreadyDone := atomic.LoadInt32(&donePtr) == 1
		oldState := atomic.LoadPointer(statePtr)

		if alreadyDone {
			return nil
		}

		if atomic.CompareAndSwapPointer(statePtr, oldState, unsafe.Pointer(&newState)) {
			atomic.StoreInt32(&donePtr, 1)
			return (*callback[T])(oldState) // note: this is illegal for nilCallback
		}
	}

	panic("unreachable")
}

func Create[T any]() Promise[T] {
	var done int32
	var state unsafe.Pointer

	atomic.StoreInt32(&done, 0)
	atomic.StorePointer(&state, unsafe.Pointer(nilCallback))

	p := new(prom[T])

	p.doneFunc = func() bool { return *((*attempt.Attempt[T])(atomic.LoadPointer(&state))) != nil }

	p.valueFunc = func() option.Option[attempt.Attempt[T]] {
		if res := *((*attempt.Attempt[T])(atomic.LoadPointer(&state))); res != nil {
			return option.Some(res)
		}

		return option.None[attempt.Attempt[T]]()
	}

	p.onCompleteFunc = func(f func(attempt.Attempt[T])) future.Future[T] {
		for i := 0; i < 1000; i++ {
			oldState := atomic.LoadPointer(&state)

			if res := *((*attempt.Attempt[T])(oldState)); res != nil {
				fmt.Println(res)
				fmt.Printf("nope %p\n", f)
				f(res)
				return p.Future()
			}

			fmt.Printf("add %p\n", f)

			next := (*callback[T])(oldState)
			newState := &callback[T]{f: f, next: next}

			if atomic.CompareAndSwapPointer(&state, oldState, unsafe.Pointer(&newState)) {
				return p.Future()
			}
		}

		panic("unreachable")
	}

	p.tryCompleteFunc = func(a attempt.Attempt[T]) bool {
		callbacks := swapCallbacksForValue(&state, a)

		switch {
		case callbacks == nil:
			// already completed
			return false
		case unsafe.Pointer(callbacks) == unsafe.Pointer(nilCallback):
			// successfully completed without listeners
			return true
		default:
			// successfully completed with listeners
			cb := reverseCallbackListAndRemoveNil[T](callbacks)
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
