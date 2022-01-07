package option

func Some[T any](value T) Option[T] { return &some[T]{value: value} }

var _ = Option[any](&some[any]{})

type some[T any] struct {
	value T
}

func (s *some[T]) Must() T                        { return s.value }
func (s *some[T]) Or(alt T) T                     { return s.Must() }
func (s *some[T]) OrErr(err error) (T, error)     { return s.Must(), nil }
func (s *some[T]) Get() (T, bool)                 { return s.value, true }
func (s *some[T]) OrElse(alt Option[T]) Option[T] { return s }
func (s *some[T]) Empty() bool                    { return false }
func (s *some[T]) NonEmpty() bool                 { return true }
func (s *some[T]) Then(f func(T))                 { f(s.Must()) }

func (s *some[T]) Filter(pred func(T) bool) Option[T] {
	if pred(s.Must()) {
		return s
	}

	return None[T]()
}
