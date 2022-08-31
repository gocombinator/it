package it

// ClosedRange creates iterator from closed range using step 1.
func ClosedRange[T Number](start, end T) Iterator[T] {
	return StepClosedRange(start, end, 1)
}
