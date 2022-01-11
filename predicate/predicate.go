package predicate

import "github.com/joa/go18beta/function"

type Predicate[T any] interface{ function.Function[T, bool] }

// always true

type alwaysTrue[T any] struct{}

func (p *alwaysTrue[T]) Apply(T) bool { return true }

func AlwaysTrue[T any]() Predicate[T] { return &alwaysTrue[T]{} }

// always false

type alwaysFalse[T any] struct{}

func (p *alwaysFalse[T]) Apply(T) bool { return false }

func AlwaysFalse[T any]() Predicate[T] { return &alwaysFalse[T]{} }

// not

type not[T any] struct{ p Predicate[T] }

func (p *not[T]) Apply(in T) bool { return !p.p.Apply(in) }

func Not[T any](p Predicate[T]) Predicate[T] { return &not[T]{p} }

// or

type or[T any] struct{ p, q Predicate[T] }

func (p *or[T]) Apply(in T) bool { return p.p.Apply(in) || p.q.Apply(in) }

func Or[T any](p, q Predicate[T]) Predicate[T] { return &or[T]{p, q} }

// and

type and[T any] struct{ p, q Predicate[T] }

func (p *and[T]) Apply(in T) bool { return p.p.Apply(in) && p.q.Apply(in) }

func And[T any](p, q Predicate[T]) Predicate[T] { return &and[T]{p, q} }

// not nil

type notNil[T comparable] struct{}

func (p *notNil[T]) Apply(in T) bool {
	var zero T
	return in != zero
}

func NotNil[T comparable]() Predicate[T] { return &notNil[T]{} }
