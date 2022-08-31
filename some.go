package it

// Some yields true if at least one value satisfies predicate, false otherwise.
func Some[T any](predicate func(T) bool) func(Iterator[T]) Iterator[bool] {
	return func(it Iterator[T]) Iterator[bool] {
		for {
			if yield := it.Next(); !yield.Done {
				if predicate(yield.Value) {
					return OfValues(true)
				}
			} else {
				return OfValues(false)
			}
		}
	}
}
