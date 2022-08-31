package it

// Window yields width sized value windows.
// Returns empty iterator if number of values is less than window width.
// Yielded window arrays have length and capacity equal to width.
// See [UnsafeWindow] for alternative strategy that avoids allocations.
func Window[T any](width int) func(Iterator[T]) Iterator[[]T] {
	return func(it Iterator[T]) Iterator[[]T] {
		return Pipe2(
			it,
			UnsafeWindow[T](width),
			Map(func(value UnsafeWindowValue[T]) []T {
				var values = make([]T, width)
				var a = value.Offset
				var b = width - a
				copy(values, value.Values[a:width])
				copy(values[b:], value.Values[0:a])
				return values
			}),
		)
	}
}
