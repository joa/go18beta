package attempt

type Attempt[T any] interface {
	// Success is true when this was a successful attempt.
	Success() bool

	// Failure is true when this was a failed attempt.
	Failure() bool

	// Get the value of the attempt.
	// Panics with FailureReason if the attempt failed.
	Get() T

	// Err for a failed attempt.
	// nil if the attempt was successful.
	Err() error
}
