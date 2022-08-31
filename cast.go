package it

// Cast performs map-like conversion between provided numeric types.
func Cast[T, U Number](it Iterator[T]) Iterator[U] {
	return Map(func(value T) U { return U(value) })(it)
}
