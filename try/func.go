package try

func Map[T, U any](a Try[T], f func(T) U) Try[U] {
	return FlatMap(a, func(x T) Try[U] {
		return Success(f(x))
	})
}

func FlatMap[T, U any](a Try[T], f func(T) Try[U]) Try[U] {
	if a.Failure() {
		// TODO: if T == U we can get rid of allocation
		return Failure[U](a.Err())
	}

	return f(a.Must())
}
