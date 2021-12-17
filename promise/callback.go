package promise

import (
	"sync/atomic"

	"github.com/joa/go18beta/attempt"
)

type callback[T any] struct {
	f     func(a attempt.Attempt[T])
	next  *callback[T]
	value atomic.Value
}

func (cb *callback[T]) run() {
	cb.f(cb.value.Load().(attempt.Attempt[T]))
	cb.value.Store(attempt.Attempt[T](nil))
	cb.next = nil
}

func (cb *callback[T]) dispatch(value attempt.Attempt[T]) {
	cb.value.Store(value)
	go cb.run()
}

var nilCallback = any(&callback[any]{})

func reverseCallbackListAndRemoveNil[T any](callbacks *callback[T]) *callback[T] {
	var (
		current  = callbacks
		previous *callback[T]
		next     *callback[T]
	)

	for current != nilCallback && current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	return previous
}
