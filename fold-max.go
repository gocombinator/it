package it

// FoldMax folds values using greater-than operator yielding single value.
func FoldMax[T Number](initial T) func(Iterator[T]) Iterator[T] {
	return Fold(func(a, b T) T {
		if a > b {
			return a
		}
		return b
	}, initial)
}
