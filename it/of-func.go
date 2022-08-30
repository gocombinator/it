package it

// OfFunc creates iterator yielding values from function's results.
func OfFunc[T any](next func() T) Iterator[T] {
	return &nextIterator[T]{func() Yield[T] {
		return Yield[T]{Value: next()}
	}}
}
