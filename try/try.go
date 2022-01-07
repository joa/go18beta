package try

type Try[T any] interface {
	// Success is true when this was a successful trey.
	Success() bool

	// Failure is true when this was a failed try.
	Failure() bool

	// Must return the value of the try.
	// Panics with FailureReason in case of failure.
	Must() T

	// Get the value.
	// Returns the value and true in case of success.
	// Returns the zero value for T and false in case of failure.
	Get() (T, bool)

	// Err for a failed try.
	// nil in case of success.
	Err() error
}
