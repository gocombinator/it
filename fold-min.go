package it

// FoldMin folds values using less-than operator yielding single value.
func FoldMin[T Number](initial T) func(Iterator[T]) Iterator[T] {
	return Fold(func(a, b T) T {
		if a < b {
			return a
		}
		return b
	}, initial)
}
