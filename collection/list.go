package collection

type List[E comparable] interface {
	Collection[E]

	AddAt(i int, e E)

	GetAt(i int) E

	IndexOf(e E) int

	RemoveAt(i int) E

	SetAt(i int, e E) E
}
