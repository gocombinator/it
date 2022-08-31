package it

// Add folds values using initial value 0 and addition binary operator yielding single value.
func Add[T Number](it Iterator[T]) Iterator[T] {
	return Fold(func(a, b T) T { return a + b }, 0)(it)
}
