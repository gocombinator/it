package it

// Take yields up to n values.
func Take[T any](n int) func(Iterator[T]) Iterator[T] {
	return func(in Iterator[T]) Iterator[T] {
		var i = 0
		return OfNext(func() Yield[T] {
			if i < n {
				i++
				return in.Next()
			}
			return Yield[T]{Done: true}
		})
	}
}
