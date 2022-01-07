package future

import (
	"context"
	"time"

	"github.com/joa/go18beta/option"
	"github.com/joa/go18beta/try"
)

type Future[T any] interface {
	// Done is true if a value is present.
	Done() bool

	// Value of the future.
	//
	// Absent while the future hasn't been completed.
	Value() option.Option[try.Try[T]]

	// FallbackTo anothr feature in case this one fails.
	FallbackTo(f Future[T]) Future[T]

	// FailAfter a specified amount of time with ErrTimeout.
	FailAfter(d time.Duration) Future[T]

	// Then - Call f when the future completes successfully.
	//
	// f is called in its own go routine.
	Then(f func(value T)) Future[T]

	// Catch - Call f when the future fails.
	//
	// f is called in its own go routine.
	Catch(f func(err error)) Future[T]

	// Recover an error with a value.
	//
	// Panics within f are automatically propagate as try.Failure.
	// f is NOT called in its own go routine.
	Recover(f func(err error) T) Future[T]

	// FlatRecover an error with a new future.
	//
	// Panics within f are automatically propagate as try.Failure.
	// f is NOT called in its own go routine.
	FlatRecover(f func(err error) Future[T]) Future[T]

	// OnComplete - Call f once this future completes.
	//
	// f is called in its own go routine.
	OnComplete(f func(try.Try[T])) Future[T]

	// Chan - Convert this future into a channel.
	//
	// The channel is written exactly once the future completes.
	// It is closed afterwards.
	Chan() <-chan try.Try[T]
}

// Go - Call f in a go routine and complete the future with its value.
func Go[T any](f func() (T, error)) Future[T] {
	p := Create[T]()
	go func() {
		// c.f. https://go.dev/ref/spec#Go_statements
		// > The function value and parameters are evaluated as usual in the calling goroutine
		p.Complete(try.Func(f))
	}()
	return p.Future()
}

// Chan to Future conversion.
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
