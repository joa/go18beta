package future

import (
	"time"

	"github.com/joa/go18beta/attempt"
	"github.com/joa/go18beta/option"
)

type Future[T any] interface {
	Done() bool

	Value() option.Option[attempt.Attempt[T]]

	FallbackTo(f Future[T]) Future[T]

	FailAfter(d time.Duration) Future[T]

	Then(f func(value T)) Future[T]

	Catch(f func(err error)) Future[T]

	Recover(f func(err error) T) Future[T]

	FlatRecover(f func(err error) Future[T]) Future[T]

	OnComplete(f func(attempt.Attempt[T])) Future[T]
}
