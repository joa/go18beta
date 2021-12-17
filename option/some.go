package option

func Some[T any](value T) Option[T] { return &some[T]{value: value} }

var _ = Option[any](&some[any]{})

type some[T any] struct {
	value T
}

func (s *some[T]) Get() T                         { return s.value }
func (s *some[T]) GetOrElse(alt T) T              { return s.Get() }
func (s *some[T]) GetOrErr(err error) (T, error)  { return s.Get(), nil }
func (s *some[T]) OrElse(alt Option[T]) Option[T] { return s }
func (s *some[T]) Empty() bool                    { return false }
func (s *some[T]) NonEmpty() bool                 { return true }
func (s *some[T]) Then(f func(T))                 { f(s.Get()) }

func (s *some[T]) Filter(pred func(T) bool) Option[T] {
	if pred(s.Get()) {
		return s
	}

	return None[T]()
}
