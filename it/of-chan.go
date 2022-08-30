package it

// OfChan creates iterator from channel.
func OfChan[T any](in <-chan T) Iterator[T] {
	return OfNext(func() Yield[T] {
		if v, ok := <-in; ok {
			return Yield[T]{Value: v}
		}
		return Yield[T]{Done: true}
	})
}
