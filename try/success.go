package try

func Success[T any](value T) Try[T] { return &successfulTry[T]{value: value} }

var _ = Try[any](&successfulTry[any]{})

type successfulTry[T any] struct{ value T }

func (sa successfulTry[T]) Success() bool  { return true }
func (sa successfulTry[T]) Failure() bool  { return false }
func (sa successfulTry[T]) Must() T        { return sa.value }
func (sa successfulTry[T]) Get() (T, bool) { return sa.value, true }
func (sa successfulTry[T]) Err() error     { return nil }
