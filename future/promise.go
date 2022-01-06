package future

import (
	"github.com/joa/go18beta/attempt"
)

type Promise[T any] interface {
	TryComplete(a attempt.Attempt[T]) bool

	Complete(a attempt.Attempt[T]) Promise[T]

	CompleteWith(f Future[T]) Promise[T]

	Future() Future[T]

	Failure(err error) Promise[T]

	Success(res T) Promise[T]
}

func Success[T any](value T) Promise[T] { return ValueOf(attempt.Success(value)) }

func Failure[T any](err error) Promise[T] { return ValueOf(attempt.Failure[T](err)) }

func ValueOf[T any](value attempt.Attempt[T]) Promise[T] { return kept(value) }