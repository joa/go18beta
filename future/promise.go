package future

import "github.com/joa/go18beta/try"

type Promise[T any] interface {
	// TryComplete this promise.
	//
	// Returns true if the promise was completed; false otherwise.
	TryComplete(a try.Try[T]) bool

	// MustComplete the promise.
	//
	// Panics if the promise has already been completed.
	MustComplete(a try.Try[T]) Promise[T]

	// CompleteWith another future.
	CompleteWith(f Future[T]) Promise[T]

	// Future of this promise.
	Future() Future[T]

	// Reject the promise.
	//
	// This method completes a promise.
	//
	// Panics if the promise has already been completed.
	Reject(err error) Promise[T]

	// Resolve the promise.
	//
	// This method completes a promise.
	//
	// Panics if the promise has already been completed.
	Resolve(res T) Promise[T]
}

// PromiseOf a known Try.
func PromiseOf[T any](value try.Try[T]) Promise[T] { return kept(value) }
