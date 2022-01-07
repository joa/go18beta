package future

import (
	"github.com/joa/go18beta/pair"
	"github.com/joa/go18beta/try"
)

func Map[T, U any](f Future[T], m func(T) U) Future[U] {
	p := Create[U]()
	f.OnComplete(func(value try.Try[T]) { p.MustComplete(try.Map(value, m)) })
	return p.Future()
}

func FlatMap[T, U any](f Future[T], m func(value T) Future[U]) Future[U] {
	p := Create[U]()
	f.OnComplete(func(value try.Try[T]) {
		mapped := try.Map[T, Future[U]](value, m)

		if mapped.Success() {
			p.CompleteWith(mapped.Must())
		} else {
			p.MustComplete(try.Failure[U](mapped.Err()))
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

	x.OnComplete(func(xa try.Try[T]) {
		if xa.Failure() {
			p.MustComplete(try.Failure[pair.Pair[T, U]](xa.Err()))
		} else {
			y.OnComplete(func(ya try.Try[U]) {
				if ya.Failure() {
					p.MustComplete(try.Failure[pair.Pair[T, U]](ya.Err()))
				} else {
					p.Resolve(pair.Pair[T, U]{
						X: xa.Must(),
						Y: ya.Must(),
					})
				}
			})
		}
	})

	return p.Future()
}
