package try

func Success[T any](value T) Try[T] { return &trySucc[T]{value: value} }

var _ = Try[any](&trySucc[any]{})

type trySucc[T any] struct{ value T }

func (ts trySucc[T]) Success() bool                         { return true }
func (ts trySucc[T]) Failure() bool                         { return false }
func (ts trySucc[T]) Must() T                               { return ts.value }
func (ts trySucc[T]) Or(T) T                                { return ts.Must() }
func (ts trySucc[T]) Get() (T, error)                       { return ts.Must(), ts.Err() }
func (ts trySucc[T]) Err() error                            { return nil }
func (ts trySucc[T]) Recover(func(error) T) Try[T]          { return ts }
func (ts trySucc[T]) FlatRecover(func(error) Try[T]) Try[T] { return ts }
func (ts trySucc[T]) OrElse(Try[T]) Try[T]                  { return ts }
func (ts trySucc[T]) Fold(f func(T), _ func(error))         { f(ts.Must()) }
