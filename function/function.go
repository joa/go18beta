package function

import "fmt"

type Function[P0, R any] interface{ Apply(p0 P0) R }

type Procedure[R any] interface{ Apply() R }

type Function1[P0, R any] interface{ Apply(p0 P0) R } // TODO:  generic type cannot be alias
type Function2[P0, P1, R any] interface{ Apply(p0 P0, p1 P1) R }
type Function3[P0, P1, P2, R any] interface{ Apply(p0 P0, p1 P1, p2 P2) R }
type Function4[P0, P1, P2, P3, R any] interface {
	Apply(p0 P0, p1 P1, p2 P2, p3 P3) R
}
type Function5[P0, P1, P2, P3, P4, R any] interface {
	Apply(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) R
}
type Function6[P0, P1, P2, P3, P4, P5, R any] interface {
	Apply(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) R
}
type Function7[P0, P1, P2, P3, P4, P5, P6, R any] interface {
	Apply(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R
}
type Function8[P0, P1, P2, P3, P4, P5, P6, P7, R any] interface {
	Apply(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R
}

// to string

type stringFunc[T fmt.Stringer] struct{}

func (sf *stringFunc[T]) Apply(in T) string { return in.String() }

func String[T fmt.Stringer]() Function[T, string] {
	return &stringFunc[T]{}
}

// identity

type identFunc[T any] struct{}

func (id *identFunc[T]) Apply(in T) T { return in }

func Identity[T any]() Function[T, T] {
	return &identFunc[T]{}
}
