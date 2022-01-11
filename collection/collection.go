package collection

type Collection[E comparable] interface {
	Iterable[E]

	Add(e E) bool

	Clear()

	Contains(e E) bool

	IsEmpty() bool

	Remove(e E) bool

	Size() int

	ToSlice() []E
}
