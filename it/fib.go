package it

// Fib yields Fibonacci sequence.
func Fib[T Number]() Iterator[T] {
	var a, b T = 1, 1
	return OfFunc(func() T {
		a, b = b, a+b
		return a
	})
}
