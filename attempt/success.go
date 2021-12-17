package attempt

func Success[T any](value T) Attempt[T] { return &successfulAttempt[T]{value: value} }

var _ = Attempt[any](&successfulAttempt[any]{})

type successfulAttempt[T any] struct{ value T }

func (sa successfulAttempt[T]) Success() bool { return true }
func (sa successfulAttempt[T]) Failure() bool { return false }
func (sa successfulAttempt[T]) Get() T        { return sa.value }
func (sa successfulAttempt[T]) Err() error    { return nil }
