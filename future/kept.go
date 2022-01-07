package future

import (
	"github.com/joa/go18beta/option"
	"github.com/joa/go18beta/try"
)

func alwaysDone() bool { return true }

func kept[T any](a try.Try[T]) Promise[T] {
	res := option.Some(a)

	p := new(prom[T])

	p.doneFunc = alwaysDone

	p.valueFunc = func() option.Option[try.Try[T]] { return res }

	p.onCompleteFunc = func(f func(try.Try[T])) Future[T] {
		go f(a)
		return p.Future()
	}

	p.tryCompleteFunc = func(_ try.Try[T]) bool { return false }

	return p
}
