package it

// Reduce performs fold using first value as initial value.
// Returns empty iterator if input iterator is empty.
// See [Fold] for alternative reduction strategy.
func Reduce[T any](f func(T, T) T) func(Iterator[T]) Iterator[T] {
	return func(it Iterator[T]) Iterator[T] {
		if yield := it.Next(); !yield.Done {
			return Fold(f, yield.Value)(it)
		} else {
			return Empty[T]()
		}
	}
}
