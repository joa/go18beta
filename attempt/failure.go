package attempt

func Failure[T any](err error) Attempt[T] { return &failedAttempt[T]{err: err} }

var _ = Attempt[any](&failedAttempt[any]{})

type failedAttempt[T any] struct{ err error }

func (sa failedAttempt[T]) Success() bool        { return false }
func (sa failedAttempt[T]) Failure() bool        { return true }
func (sa failedAttempt[T]) Get() T               { panic(sa.FailureReason()) }
func (sa failedAttempt[T]) FailureReason() error { return sa.err }
