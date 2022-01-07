package future

import "github.com/joa/go18beta/try"

type Promise[T any] interface {
	TryComplete(a try.Try[T]) bool

	Complete(a try.Try[T]) Promise[T]

	CompleteWith(f Future[T]) Promise[T]

	Future() Future[T]

	Failure(err error) Promise[T]

	Success(res T) Promise[T]
}

func Success[T any](value T) Promise[T] { return ValueOf(try.Success(value)) }

func Failure[T any](err error) Promise[T] { return ValueOf(try.Failure[T](err)) }

func ValueOf[T any](value try.Try[T]) Promise[T] { return kept(value) }
