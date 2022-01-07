package try

func Failure[T any](err error) Try[T] { return &failedTry[T]{err: err} }

var _ = Try[any](&failedTry[any]{})

type failedTry[T any] struct{ err error }

func (sa failedTry[T]) Success() bool { return false }
func (sa failedTry[T]) Failure() bool { return true }
func (sa failedTry[T]) Must() T       { panic(sa.Err()) }
func (sa failedTry[T]) Or(alt T) T    { return alt }
func (sa failedTry[T]) Get() (T, error) {
	var zero T
	return zero, sa.Err()
}
func (sa failedTry[T]) Err() error { return sa.err }
