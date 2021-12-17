package promise

import (
	"github.com/joa/go18beta/attempt"
	"github.com/joa/go18beta/future"
)

type Promise[T any] interface {
	TryComplete(a attempt.Attempt[T]) bool

	Complete(a attempt.Attempt[T]) Promise[T]

	CompleteWith(f future.Future[T]) Promise[T]

	Future() future.Future[T]

	Failure(err error) Promise[T]

	Success(res T) Promise[T]
}

func Create[T any]() Promise[T] {
	var todo Promise[T]
	return todo
}

func Success[T any](v T) future.Future[T] {
	return attempt.Success(v)
}
