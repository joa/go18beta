package option

type Option[T any] interface {
	Must() T

	Or(alt T) T

	Get() (T, bool)

	OrErr(err error) (T, error)

	OrElse(alt Option[T]) Option[T]

	Empty() bool

	NonEmpty() bool

	Filter(pred func(T) bool) Option[T]

	Then(f func(T))
}
