package attempt

func Map[T, U any](a Attempt[T], f func(T) U) Attempt[U] {
	return FlatMap(a, func(x T) Attempt[U] {
		return Success(f(x))
	})
}

func FlatMap[T, U any](a Attempt[T], f func(T) Attempt[U]) Attempt[U] {
	if a.Failure() {
		var hack = any(a)
		return hack.(*failedAttempt[U])
	}

	return f(a.Get())
}
