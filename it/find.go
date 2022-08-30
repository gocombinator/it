package it

// Find yields first value of iterator that satisfies predicate.
// Returns empty iterator if no value satisfies predicate.
func Find[T any](predicate func(T) bool) func(Iterator[T]) Iterator[T] {
	return func(it Iterator[T]) Iterator[T] {
		for {
			if yield := it.Next(); !yield.Done {
				if predicate(yield.Value) {
					return OfValues(yield.Value)
				}
			} else {
				return Empty[T]()
			}
		}
	}
}
