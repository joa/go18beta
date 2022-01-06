package future

import (
	"github.com/joa/go18beta/attempt"
	"github.com/joa/go18beta/option"
)

func alwaysDone() bool { return true }

func kept[T any](a attempt.Attempt[T]) Promise[T] {
	res := option.Some(a)

	p := new(prom[T])

	p.doneFunc = alwaysDone

	p.valueFunc = func() option.Option[attempt.Attempt[T]] { return res }

	p.onCompleteFunc = func(f func(attempt.Attempt[T])) Future[T] {
		go f(a)
		return p.Future()
	}

	p.tryCompleteFunc = func(_ attempt.Attempt[T]) bool { return false }

	return p
}
