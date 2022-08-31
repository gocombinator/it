package it

// Combinations yields k-combinations of iterator values.
// See [UnsafeCombinations] for faster version.
func Combinations[T any](k int) func(Iterator[T]) Iterator[[]T] {
	return func(it Iterator[T]) Iterator[[]T] {
		return Pipe2(
			it,
			UnsafeCombinations[T](k),
			Map(func(values []T) []T {
				var valuesCopy = make([]T, len(values))
				copy(valuesCopy, values)
				return valuesCopy
			}),
		)
	}
}
