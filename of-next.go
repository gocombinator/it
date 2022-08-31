package it

// OfNext creates iterator from next function.
func OfNext[T any](next func() Yield[T]) Iterator[T] {
	return &nextIterator[T]{next}
}
