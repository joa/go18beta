package option

func Map[T, U any](a Option[T], f func(T) U) Option[U] {
	return FlatMap(a, func(x T) Option[U] {
		return Some(f(x))
	})
}

func FlatMap[T, U any](a Option[T], f func(T) Option[U]) Option[U] {
	if a.Empty() {
		return None[U]()
	}

	return f(a.Get())
}
