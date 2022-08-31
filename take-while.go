package it

// TakeWhile yields values until predicate is not satisfied.
func TakeWhile[T any](predicate func(T) bool) func(Iterator[T]) Iterator[T] {
	return func(in Iterator[T]) Iterator[T] {
		var done bool
		return OfNext(func() Yield[T] {
			if done {
				return Yield[T]{Done: true}
			}
			yield := in.Next()
			if yield.Done {
				done = true
				return Yield[T]{Done: true}
			}
			if !predicate(yield.Value) {
				done = true
				return Yield[T]{Done: true}
			}
			return yield
		})
	}
}
