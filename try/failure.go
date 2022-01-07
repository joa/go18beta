package try

func Failure[T any](err error) Try[T] { return &tryErr[T]{err: err} }

var _ = Try[any](&tryErr[any]{})

type tryErr[T any] struct{ err error }

func (te tryErr[T]) Success() bool { return false }
func (te tryErr[T]) Failure() bool { return true }
func (te tryErr[T]) Must() T       { panic(te.Err()) }
func (te tryErr[T]) Or(alt T) T    { return alt }
func (te tryErr[T]) Get() (T, error) {
	var zero T
	return zero, te.Err()
}
func (te tryErr[T]) Err() error               { return te.err }
func (te tryErr[T]) OrElse(alt Try[T]) Try[T] { return alt }
func (te tryErr[T]) Recover(f func(err error) T) (res Try[T]) {
	return te.FlatRecover(func(err error) Try[T] {
		return Success(f(err))
	})
}
func (te tryErr[T]) FlatRecover(f func(err error) Try[T]) (res Try[T]) {
	defer panicToFailure(&res)
	res = f(te.Err())
	return
}
