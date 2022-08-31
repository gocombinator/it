package it

// Fold performs reduction using initial value with arbitrary return type yielding single value.
// See [Reduce] for alternative reduction strategy.
func Fold[T any, A any](f func(A, T) A, initial A) func(Iterator[T]) Iterator[A] {
	return func(it Iterator[T]) Iterator[A] {
		var value = initial
		for {
			if yield := it.Next(); !yield.Done {
				value = f(value, yield.Value)
			} else {
				return OfValues(value)
			}
		}
	}
}
