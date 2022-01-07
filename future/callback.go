package future

import (
	"sync/atomic"
	"unsafe"

	"github.com/joa/go18beta/try"
)

type callback[T any] struct {
	f     func(a try.Try[T])
	next  *callback[T]
	value *atomic.Value
}

func (cb *callback[T]) run() {
	cb.f(cb.value.Load().(try.Try[T]))
	cb.value = nil
	cb.next = nil
}

func (cb *callback[T]) dispatch(value try.Try[T]) {
	cb.value.Store(value)
	go cb.run()
}

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
