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

// vtable is a list of a promise's virtual functions.
type vtable[T any] struct {
	doneFunc        func() bool
	valueFunc       func() option.Option[attempt.Attempt[T]]
	onCompleteFunc  func(func(attempt.Attempt[T])) future.Future[T]
	tryCompleteFunc func(a attempt.Attempt[T]) bool
}

// prom is an abstract promise.
//
// instances must provide a vtable for implementations of the abstract methods.
type prom[T any] struct {
	vtable[T]
}

// future

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
		if a.Success() {
			q.Complete(a)
		} else {
			q.CompleteWith(f(a.Err()))
		}
	})

	return q.Future()
}

func (p *prom[T]) Done() bool { return p.doneFunc() }

func (p *prom[T]) Value() option.Option[attempt.Attempt[T]] { return p.valueFunc() }

func (p *prom[T]) OnComplete(f func(attempt.Attempt[T])) future.Future[T] {
	return p.onCompleteFunc(f)
}

// promise

func (p *prom[T]) TryComplete(a attempt.Attempt[T]) bool {
	return p.tryCompleteFunc(a)
}

func (p *prom[T]) Complete(a attempt.Attempt[T]) Promise[T] {
	if p.TryComplete(a) {
		return p
	} else {
		panic("promise already completed")
	}
}

func (p *prom[T]) CompleteWith(f future.Future[T]) Promise[T] {
	f.OnComplete(func(a attempt.Attempt[T]) {
		p.TryComplete(a)
	})

	return p
}

func (p *prom[T]) Future() future.Future[T] {
	return future.Future[T](p)
}

func (p *prom[T]) Failure(err error) Promise[T] {
	return p.Complete(attempt.Failure[T](err))
}

func (p *prom[T]) Success(res T) Promise[T] {
	return p.Complete(attempt.Success[T](res))
}
