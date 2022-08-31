package it

// Skip drops n values from iterator.
func Skip[T any](n int) func(Iterator[T]) Iterator[T] {
	return func(in Iterator[T]) Iterator[T] {
		for i := 0; i < n; i++ {
			if yield := in.Next(); yield.Done {
				break
			}
		}
		return in
	}
}
