package it

// Last yields last value from non empty iterator.
// Returns empty iterator if iterator is empty.
func Last[T any](it Iterator[T]) Iterator[T] {
	var result = Yield[T]{Done: true}
	for {
		if yield := it.Next(); !yield.Done {
			result = yield
		} else if result.Done {
			return Empty[T]()
		} else {
			return OfValues(result.Value)
		}
	}
}
