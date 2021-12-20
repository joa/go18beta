package promise

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func updateState[O, N any](state *unsafe.Pointer, old *O, new *N) bool {
	return atomic.CompareAndSwapPointer(state,
		unsafe.Pointer(old),
		unsafe.Pointer(new))
}
func Create[T any]() Promise[T] {
	var state unsafe.Pointer

	getState := func() unsafe.Pointer {
		return atomic.LoadPointer(&state)
	}

	cb := nilCallback.(*callback[any])
	atomic.StorePointer(&state, unsafe.Pointer(cb))

	fmt.Println((*callback[any])(getState()))

	x := new(interface{})

	fmt.Println(updateState(&state, cb, x))

	fmt.Println(getState())

	var todo Promise[T]
	return todo
}
