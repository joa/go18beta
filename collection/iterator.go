package collection

type Iterator[E comparable] interface {
	HasNext() bool
	Next() E
}

type Iterable[E comparable] interface {
	Iterator() Iterator[E]
}
