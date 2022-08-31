package it

type UnsafeWindowValue[T any] struct {
	Values []T
	Offset int
}

// UnsafeWindow yields width sized value windows.
// Yields shared ring array with offset.
// Yielded values are readonly and have to be consumed before next iteration.
// See [Window] for alternative strategy.
func UnsafeWindow[T any](width int) func(Iterator[T]) Iterator[UnsafeWindowValue[T]] {
	return func(it Iterator[T]) Iterator[UnsafeWindowValue[T]] {

		// Using single window allocation.
		var values = make([]T, width)

		// Maybe prefill window.
		var offset = 0
		for {
			if yield := it.Next(); !yield.Done {
				values[offset] = yield.Value
				offset++
				if offset == width-1 {
					break
				}
			} else {
				break
			}
		}

		return OfNext(func() Yield[UnsafeWindowValue[T]] {
			if yield := it.Next(); !yield.Done {
				values[offset] = yield.Value
				offset++
				if offset == width {
					offset = 0
				}
				return Yield[UnsafeWindowValue[T]]{
					Value: UnsafeWindowValue[T]{
						Values: values,
						Offset: offset,
					},
				}
			} else {
				return Yield[UnsafeWindowValue[T]]{Done: true}
			}
		})
	}
}
