package try

import (
	"errors"
	"fmt"
)

func Map[T, U any](a Try[T], f func(T) U) Try[U] {
	return FlatMap(a, func(x T) Try[U] {
		return Success(f(x))
	})
}

func FlatMap[T, U any](a Try[T], f func(T) Try[U]) (res Try[U]) {
	if a.Failure() {
		// TODO: if T == U we can get rid of allocation
		return Failure[U](a.Err())
	}

	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case error:
				res = Failure[U](r)
			case string:
				res = Failure[U](errors.New(r))
			default:
				res = Failure[U](fmt.Errorf("%v", r))
			}
		}
	}()

	res = f(a.Must())

	return
}
