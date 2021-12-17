package option

func None[T any]() Option[T] {
	//TODO: check if we can get rid of this alloc
	return &none[T]{}
}

var _ = Option[any](&none[any]{})

type none[T any] struct{}

func (n *none[T]) Get() T            { panic(ErrEmpty) }
func (n *none[T]) GetOrElse(alt T) T { return alt }
func (n *none[T]) GetOrErr(err error) (T, error) {
	var zero T
	return zero, err
}
func (n *none[T]) OrElse(alt Option[T]) Option[T]     { return alt }
func (n *none[T]) Empty() bool                        { return true }
func (n *none[T]) NonEmpty() bool                     { return false }
func (n *none[T]) Filter(pred func(T) bool) Option[T] { return n }
func (n *none[T]) Then(f func(T))                     {}
