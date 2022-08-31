package it

// StepClosedRange creates iterator yielding values in closed range using step.
func StepClosedRange[T Number](start, end, step T) Iterator[T] {
	var i = start
	return OfNext(func() Yield[T] {
		if i > end {
			return Yield[T]{Done: true}
		}
		var value = i
		i += step
		return Yield[T]{Value: value}
	})
}
