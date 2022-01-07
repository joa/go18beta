package future

import (
	"time"

	"github.com/joa/go18beta/option"
	"github.com/joa/go18beta/try"
)

type Future[T any] interface {
	Done() bool

	Value() option.Option[try.Try[T]]

	FallbackTo(f Future[T]) Future[T]

	FailAfter(d time.Duration) Future[T]

	Then(f func(value T)) Future[T]

	Catch(f func(err error)) Future[T]

	Recover(f func(err error) T) Future[T]

	FlatRecover(f func(err error) Future[T]) Future[T]

	OnComplete(f func(try.Try[T])) Future[T]
}
