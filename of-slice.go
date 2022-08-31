package it

// OfSlice creates iterator from slice.
func OfSlice[T any](in []T) Iterator[T] {
	var i = 0
	var value T
	return OfNext(func() Yield[T] {
		if i < len(in) {
			value = in[i]
			i++
			return Yield[T]{Value: value}
		}
		return Yield[T]{Done: true}
	})
}
