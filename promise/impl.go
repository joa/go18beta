package promise

import (
	"time"

	"github.com/joa/go18beta/attempt"
	"github.com/joa/go18beta/future"
	"github.com/joa/go18beta/option"
)

var (
	_ = future.Future[any](&prom[any]{})
	_ = Promise[any](&prom[any]{})
)

type prom[T any] struct {
}

// future

func (p *prom[T]) Done() bool {

}

func (p *prom[T]) Value() option.Option[attempt.Attempt[T]] {

}

func (p *prom[T]) FallbackTo(f future.Future[T]) future.Future[T] {
	q := Create[T]()

	p.OnComplete(func(a attempt.Attempt[T]) {
		if a.Success() {
			q.Complete(a)
		} else {
			q.CompleteWith(f)
		}
	})

	return q.Future()
}

func (p *prom[T]) FailAfter(d time.Duration) future.Future[T] {
	q := Create[T]()

	p.OnComplete(func(a attempt.Attempt[T]) {
		q.TryComplete(a)
	})

	time.AfterFunc(d, func() {
		q.TryComplete(attempt.Failure[T](ErrTimeout))
	})

	return q.Future()
}

func (p *prom[T]) Then(f func(value T)) future.Future[T] {
	return p.OnComplete(func(a attempt.Attempt[T]) {
		if a.Success() {
			f(a.Get())
		}
	})
}

func (p *prom[T]) Catch(f func(err error)) future.Future[T] {
	return p.OnComplete(func(a attempt.Attempt[T]) {
		if a.Failure() {
			f(a.Err())
		}
	})
}

func (p *prom[T]) Recover(f func(err error) T) future.Future[T] {
	q := Create[T]()

	p.OnComplete(func(a attempt.Attempt[T]) {
		if a.Success() {
			q.Complete(a)
		} else {
			q.Complete(attempt.Success(f(a.Err())))
		}
	})

	return q.Future()
}

func (p *prom[T]) FlatRecover(f func(err error) future.Future[T]) future.Future[T] {
	q := Create[T]()

	p.OnComplete(func(a attempt.Attempt[T]) {

	})

	return q.Future()
}

func (p *prom[T]) OnComplete(f func(attempt.Attempt[T])) future.Future[T] {

}
