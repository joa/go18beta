package future

import (
	"context"
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

// Go - Call f in a go routine and complete the future with its value.
func Go[T any](f func() (T, error)) Future[T] {
	p := Create[T]()
	go p.Complete(try.Func(f))
	return p.Future()
}

// Chan to Future
//
// The first value produced by the channel completes the future.
func Chan[T any](ctx context.Context, ch <-chan T) Future[T] {
	p := Create[T]()
	go func() {
		select {
		case value, ok := <-ch:
			if ok {
				p.Resolve(value)
			} else {
				p.Reject(ErrChannelClosed)
			}
			return
		case <-ctx.Done():
			p.Reject(ctx.Err())
			return
		}
	}()
	return p.Future()
}
