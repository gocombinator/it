package it

// Empty returns an empty iterator.
func Empty[T any]() Iterator[T] {
	return &emptyIterator[T]{}
}
