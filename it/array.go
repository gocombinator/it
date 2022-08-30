package it

// Array consumes iterator returning an array of values.
func Array[T any](in Iterator[T]) []T {
	var out = make([]T, 0)
	for {
		if yield := in.Next(); !yield.Done {
			out = append(out, yield.Value)
		} else {
			break
		}
	}
	return out
}
