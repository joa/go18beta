package option

type Option[T any] interface {
	Must() T

	Get() (T, bool)

	GetOrElse(alt T) T

	GetOrErr(err error) (T, error)

	OrElse(alt Option[T]) Option[T]

	Empty() bool

	NonEmpty() bool

	Filter(pred func(T) bool) Option[T]

	Then(f func(T))
}
