package future

import (
	"time"

	"github.com/joa/go18beta/option"
	"github.com/joa/go18beta/try"
)

var (
	_ = Future[any](&prom[any]{})
	_ = Promise[any](&prom[any]{})
)

// vtable is a list of a promise's virtual functions.
type vtable[T any] struct {
	doneFunc        func() bool
	valueFunc       func() option.Option[try.Try[T]]
	onCompleteFunc  func(func(try.Try[T])) Future[T]
	tryCompleteFunc func(a try.Try[T]) bool
}

// prom is an abstract promise.
//
// instances must provide a vtable for implementations of the abstract methods.
type prom[T any] struct {
	vtable[T]
}

// future

func (p *prom[T]) FallbackTo(f Future[T]) Future[T] {
	q := Create[T]()

	p.OnComplete(func(a try.Try[T]) {
		if a.Success() {
			q.Complete(a)
		} else {
			q.CompleteWith(f)
		}
	})

	return q.Future()
}

func (p *prom[T]) FailAfter(d time.Duration) Future[T] {
	q := Create[T]()

	p.OnComplete(func(a try.Try[T]) {
		q.TryComplete(a)
	})

	time.AfterFunc(d, func() {
		q.TryComplete(try.Failure[T](ErrTimeout))
	})

	return q.Future()
}

func (p *prom[T]) Then(f func(value T)) Future[T] {
	return p.OnComplete(func(a try.Try[T]) {
		if a.Success() {
			f(a.Must())
		}
	})
}

func (p *prom[T]) Catch(f func(err error)) Future[T] {
	return p.OnComplete(func(a try.Try[T]) {
		if a.Failure() {
			f(a.Err())
		}
	})
}

func (p *prom[T]) Recover(f func(err error) T) Future[T] {
	q := Create[T]()

	p.OnComplete(func(a try.Try[T]) {
		if a.Success() {
			q.Complete(a)
		} else {
			q.Complete(try.Success(f(a.Err())))
		}
	})

	return q.Future()
}

func (p *prom[T]) FlatRecover(f func(err error) Future[T]) Future[T] {
	q := Create[T]()

	p.OnComplete(func(a try.Try[T]) {
		if a.Success() {
			q.Complete(a)
		} else {
			q.CompleteWith(f(a.Err()))
		}
	})

	return q.Future()
}

func (p *prom[T]) Done() bool { return p.doneFunc() }

func (p *prom[T]) Value() option.Option[try.Try[T]] { return p.valueFunc() }

func (p *prom[T]) OnComplete(f func(try.Try[T])) Future[T] {
	return p.onCompleteFunc(f)
}

// promise

func (p *prom[T]) TryComplete(a try.Try[T]) bool {
	return p.tryCompleteFunc(a)
}

func (p *prom[T]) Complete(a try.Try[T]) Promise[T] {
	if p.TryComplete(a) {
		return p
	} else {
		panic("promise already completed")
	}
}

func (p *prom[T]) CompleteWith(f Future[T]) Promise[T] {
	f.OnComplete(func(a try.Try[T]) {
		p.TryComplete(a)
	})

	return p
}

func (p *prom[T]) Future() Future[T] {
	return Future[T](p)
}

func (p *prom[T]) Failure(err error) Promise[T] {
	return p.Complete(try.Failure[T](err))
}

func (p *prom[T]) Success(res T) Promise[T] {
	return p.Complete(try.Success[T](res))
}
