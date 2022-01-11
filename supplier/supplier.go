package supplier

import "github.com/joa/go18beta/function"

type Supplier[T any] interface {
	function.Procedure[T]
}

type funcSupplier[T any] struct{ fun func() T }

func (fs *funcSupplier[T]) Apply() T { return fs.fun() }

func Func[T any](fun func() T) Supplier[T] { return &funcSupplier[T]{fun} }

type valueSupplier[T any] struct{ value T }

func (vs *valueSupplier[T]) Apply() T { return vs.value }

func Of[T any](value T) Supplier[T] { return &valueSupplier[T]{value} }
