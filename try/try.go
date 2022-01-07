package try

import (
	"errors"
	"fmt"
)

type Try[T any] interface {
	// Success is true when this was a successful trey.
	Success() bool

	// Failure is true when this was a failed try.
	Failure() bool

	// Must return the value of the try.
	//
	// Panics with Err() in case of failure.
	Must() T

	// Or returns an alternative value in case of failure.
	Or(alt T) T

	// Get the value and error.
	//
	// The value will be the zero value of T in case of failure.
	// The error will be nil in case of success.
	Get() (T, error)

	// Err for a failed try.
	// nil in case of success.
	Err() error

	Recover(f func(err error) T) Try[T]

	FlatRecover(f func(err error) Try[T]) Try[T]

	OrElse(t Try[T]) Try[T]
}

// Func - Call f and return a Try for the result.
func Func[T any](f func() (T, error)) (res Try[T]) {
	defer panicToFailure(&res)

	if value, err := f(); err == nil {
		return Success(value)
	} else {
		return Failure[T](err)
	}
}

func panicToFailure[T any](res *Try[T]) {
	if r := recover(); r != nil {
		switch r := r.(type) {
		case error:
			*res = Failure[T](r)
		case string:
			*res = Failure[T](errors.New(r))
		default:
			*res = Failure[T](fmt.Errorf("%v", r))
		}
	}
}
