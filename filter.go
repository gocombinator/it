package it

// Filter yields values that satisfy predicate.
func Filter[T any](predicate func(T) bool) func(Iterator[T]) Iterator[T] {
	return func(in Iterator[T]) Iterator[T] {
		return OfNext(func() Yield[T] {
			for {
				var yield = in.Next()
				if yield.Done {
					break
				}
				if predicate(yield.Value) {
					return yield
				}
			}
			return Yield[T]{Done: true}
		})
	}
}
