package future

import (
	"unsafe"

	"github.com/joa/go18beta/try"
)

// callback is a single callback reference.
//
// Multiple callbacks can be stored as a single linked list via the next pointer.
type callback[T any] struct {
	f    func(a try.Try[T])
	next *callback[T]
}

// dispatch the callback.
func (cb *callback[T]) dispatch(value try.Try[T]) {
	cb.next = nil
	go cb.f(value)
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
