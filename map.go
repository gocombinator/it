package it

// Map applies mapping function to each value.
func Map[T, U any](mapping func(T) U) func(Iterator[T]) Iterator[U] {
	return func(in Iterator[T]) Iterator[U] {
		return OfNext(func() Yield[U] {
			if yield := in.Next(); !yield.Done {
				return Yield[U]{Value: mapping(yield.Value)}
			}
			return Yield[U]{Done: true}
		})
	}
}
