package it

// Every yields true if all values satisfy predicate, false otherwise.
func Every[T any](predicate func(T) bool) func(Iterator[T]) Iterator[bool] {
	return func(it Iterator[T]) Iterator[bool] {
		for {
			if yield := it.Next(); !yield.Done {
				if !predicate(yield.Value) {
					return OfValues(false)
				}
			} else {
				return OfValues(true)
			}
		}
	}
}
