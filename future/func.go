package future

import "github.com/joa/go18beta/pair"

func Map[T, U any](f Future[T], m func(T) U) Future[U] {
	var todo Future[U]
	return todo
}

func FlatMap[T, U any](f Future[T], m func(value T) Future[U]) Future[U] {
	var todo Future[U]
	return todo
}

func Flatten[T, U any](f Future[T]) Future[U] {
	var todo Future[U]
	return todo
}

func Join[T, U any](x Future[T], y Future[U]) pair.Pair[Future[T], Future[U]] {
	var todo pair.Pair[Future[T], Future[U]]
	return todo
}
