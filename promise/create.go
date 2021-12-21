package promise

import (
	"github.com/joa/go18beta/attempt"
	"github.com/joa/go18beta/future"
	"github.com/joa/go18beta/option"
	"sync/atomic"
	"unsafe"
)

func swapCallbacksForValue[T any](statePtr *unsafe.Pointer, newState attempt.Attempt[T]) *callback[T] {
	for {
		state := atomic.LoadPointer(statePtr)

		if *((*attempt.Attempt[T])(atomic.LoadPointer(&state))) != nil {
			return nil
		}

		if atomic.CompareAndSwapPointer(statePtr, state, unsafe.Pointer(&newState)) {
			return (*callback[T])(state) // note: this is illegal for nilCallback
		}
	}
}

func Create[T any]() Promise[T] {
	var state unsafe.Pointer

	nilCB := nilCallback.(*callback[any])

	atomic.StorePointer(&state, unsafe.Pointer(nilCB))

	res := &prom[T]{
		doneFunc: func() bool { return *((*attempt.Attempt[T])(atomic.LoadPointer(&state))) != nil },
		valueFunc: func() option.Option[attempt.Attempt[T]] {
			if res := *((*attempt.Attempt[T])(atomic.LoadPointer(&state))); res != nil {
				return option.Some(res)
			}

			return option.None[attempt.Attempt[T]]()
		},
		onCompleteFunc: func(func(attempt.Attempt[T])) future.Future[T] {
			var x future.Future[T]
			return x
		},
		tryCompleteFunc: func(a attempt.Attempt[T]) bool {
			switch callbacks := swapCallbacksForValue(&state, a); {
			case callbacks == nil:
				// already completed
				return false
			case unsafe.Pointer(callbacks) == unsafe.Pointer(nilCB):
				// successfully completed without listeners
				return true
			default:
				// successfully completed with listeners
				cb := reverseCallbackListAndRemoveNil[T]((*callback[T])(callbacks))
				for cb != nil {
					next := cb.next
					cb.dispatch(a)
					cb = next
				}
				return true
			}
		},
	}

	return res
}
