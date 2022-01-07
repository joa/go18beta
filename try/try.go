package try

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
}
