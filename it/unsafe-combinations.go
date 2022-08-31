package it

// UnsafeCombinations yields k-combinations of values.
// Uses single array allocation.
// Avoids allocations per value making it faster than [Combinations].
// Yielded array is shared, it is being mutated, it can be used as read only array consumed immediatelly.
// Use when performance matters and yielded values are immediatelly transformed into some form of copied value.
func UnsafeCombinations[T any](k int) func(Iterator[T]) Iterator[[]T] {
	return func(it Iterator[T]) Iterator[[]T] {
		var ch = make(chan []T)
		go func() {
			var values = Array(it)
			var n = len(values)
			var combination = make([]T, k)
			var aux func(from int, index int)
			aux = func(from, index int) {
				if index == k {
					ch <- combination
					return
				}
				for i := from; i <= n-1 && n-i >= k-index; i++ {
					combination[index] = values[i]
					aux(i+1, index+1)
				}
			}
			aux(0, 0)
			close(ch)
		}()
		return OfChan(ch)
	}
}
