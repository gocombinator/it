package it

// OfValues creates iterator from one or more arguments.
func OfValues[T any](in ...T) Iterator[T] {
	return OfSlice(in)
}
