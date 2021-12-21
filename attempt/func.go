package attempt

func Map[T, U any](a Attempt[T], f func(T) U) Attempt[U] {
	return FlatMap(a, func(x T) Attempt[U] {
		return Success(f(x))
	})
}

func FlatMap[T, U any](a Attempt[T], f func(T) Attempt[U]) Attempt[U] {
	if a.Failure() {
		// TODO: if T == U we can get rid of allocation
		return Failure[U](a.Err())
	}

	return f(a.Get())
}
