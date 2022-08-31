package it

// Mul folds values using initial value 1 and multiplication binary operator yielding single value.
func Mul[T Number](it Iterator[T]) Iterator[T] {
	return Fold(func(a, b T) T { return a * b }, 1)(it)
}
