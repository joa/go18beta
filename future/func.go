package future

import (
	"github.com/joa/go18beta/attempt"
	"github.com/joa/go18beta/pair"
)

func Map[T, U any](f Future[T], m func(T) U) Future[U] {
	p := Create[U]()
	f.OnComplete(func(value attempt.Attempt[T]) {
		if value.Success() {
			p.Complete(attempt.Success(m(value.Get())))
		} else {
			p.Complete(attempt.Failure[U](value.Err()))
		}
	})
	return p.Future()
}

func FlatMap[T, U any](f Future[T], m func(value T) Future[U]) Future[U] {
	p := Create[U]()
	f.OnComplete(func(value attempt.Attempt[T]) {
		if value.Success() {
			p.CompleteWith(m(value.Get()))
		} else {
			p.Complete(attempt.Failure[U](value.Err()))
		}
	})
	return p.Future()
}

func Flatten[T, U any](f Future[T]) Future[U] {
	// intentionally left blank for the accustomed reader
	var todo Future[U]
	return todo
}

func Join[T, U any](x Future[T], y Future[U]) Future[pair.Pair[T, U]] {
	p := Create[pair.Pair[T, U]]()

	x.OnComplete(func(xa attempt.Attempt[T]) {
		if xa.Failure() {
			p.Complete(attempt.Failure[pair.Pair[T, U]](xa.Err()))
		} else {
			y.OnComplete(func(ya attempt.Attempt[U]) {
				if ya.Failure() {
					p.Complete(attempt.Failure[pair.Pair[T, U]](ya.Err()))
				} else {
					p.Success(pair.Pair[T, U]{
						X: xa.Get(),
						Y: ya.Get(),
					})
				}
			})
		}
	})

	return p.Future()
}
