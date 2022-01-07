package try

func Map[T, U any](a Try[T], f func(T) U) Try[U] {
	return FlatMap(a, func(x T) Try[U] {
		return Success(f(x))
	})
}

func MapErr[T, U any](a Try[T], f func(T) (U, error)) Try[U] {
	return FlatMap(a, func(x T) Try[U] {
		if res, err := f(x); err == nil {
			return Success(res)
		} else {
			return Failure[U](err)
		}
	})
}

func FlatMap[T, U any](a Try[T], f func(T) Try[U]) (res Try[U]) {
	if a.Failure() {
		// TODO: if T == U we can get rid of allocation
		return Failure[U](a.Err())
	}

	defer panicToFailure(&res)

	res = f(a.Must())

	return
}
