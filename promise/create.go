package promise

func Create[T any]() Promise[T] {

	var todo Promise[T]
	return todo
}
