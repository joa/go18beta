package future

import (
	"sync/atomic"
	"unsafe"

	"github.com/joa/go18beta/try"
)

// callback is a single callback reference.
//
// Multiple callbacks can be stored as a single linked list via the next pointer.
type callback[T any] struct {
	f     func(a try.Try[T])
	next  *callback[T]
	value *atomic.Value
}

// run executes the callback method with the current value and resets the state.
//
// This method bust be called only once.
func (cb *callback[T]) run() {
	cb.f(cb.value.Load().(try.Try[T]))
	cb.value = nil
	cb.next = nil
}

// dispatch memorizes the callback argument and executes the callback method asynchronous.
func (cb *callback[T]) dispatch(value try.Try[T]) {
	cb.value.Store(value)
	go cb.run()
}

// nilCallback is the terminal of the linked list of callbacks
//
// Note: Go generics don't use type erasure, so we rely on unsafe.Pointer for the
//       reference to nilCallback in some places as it's incompatible with any T that is not any.
var nilCallback = &callback[any]{}

func reverseCallbackListAndRemoveNil[T any](callbacks *callback[T]) *callback[T] {
	var (
		current  = callbacks
		previous *callback[T]
		next     *callback[T]
	)

	for unsafe.Pointer(current) != unsafe.Pointer(nilCallback) && current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	return previous
}
