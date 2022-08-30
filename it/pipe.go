package it

// Pipe1 one step pipe operator.
func Pipe1[T any, T1 any](
	it Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
) Iterator[T1] {
	return f1(it)
}

// Pipe2 two step pipe operator.
func Pipe2[T any, T1 any, T2 any](
	it Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
	f2 func(Iterator[T1]) Iterator[T2],
) Iterator[T2] {
	return f2(f1(it))
}

// Pipe3 three step pipe operator.
func Pipe3[T any, T1 any, T2 any, T3 any](
	in Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
	f2 func(Iterator[T1]) Iterator[T2],
	f3 func(Iterator[T2]) Iterator[T3],
) Iterator[T3] {
	return f3(f2(f1(in)))
}

// Pipe4 four step pipe operator.
func Pipe4[T any, T1 any, T2 any, T3 any, T4 any](
	in Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
	f2 func(Iterator[T1]) Iterator[T2],
	f3 func(Iterator[T2]) Iterator[T3],
	f4 func(Iterator[T3]) Iterator[T4],
) Iterator[T4] {
	return f4(f3(f2(f1(in))))
}

// Pipe5 five step pipe operator.
func Pipe5[T any, T1 any, T2 any, T3 any, T4 any, T5 any](
	in Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
	f2 func(Iterator[T1]) Iterator[T2],
	f3 func(Iterator[T2]) Iterator[T3],
	f4 func(Iterator[T3]) Iterator[T4],
	f5 func(Iterator[T4]) Iterator[T5],
) Iterator[T5] {
	return f5(f4(f3(f2(f1(in)))))
}

// Pipe6 six step pipe operator.
func Pipe6[T any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](
	in Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
	f2 func(Iterator[T1]) Iterator[T2],
	f3 func(Iterator[T2]) Iterator[T3],
	f4 func(Iterator[T3]) Iterator[T4],
	f5 func(Iterator[T4]) Iterator[T5],
	f6 func(Iterator[T5]) Iterator[T6],
) Iterator[T6] {
	return f6(f5(f4(f3(f2(f1(in))))))
}

// Pipe7 seven step pipe operator.
func Pipe7[T any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](
	in Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
	f2 func(Iterator[T1]) Iterator[T2],
	f3 func(Iterator[T2]) Iterator[T3],
	f4 func(Iterator[T3]) Iterator[T4],
	f5 func(Iterator[T4]) Iterator[T5],
	f6 func(Iterator[T5]) Iterator[T6],
	f7 func(Iterator[T6]) Iterator[T7],
) Iterator[T7] {
	return f7(f6(f5(f4(f3(f2(f1(in)))))))
}

// Pipe8 eight step pipe operator.
func Pipe8[T any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](
	in Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
	f2 func(Iterator[T1]) Iterator[T2],
	f3 func(Iterator[T2]) Iterator[T3],
	f4 func(Iterator[T3]) Iterator[T4],
	f5 func(Iterator[T4]) Iterator[T5],
	f6 func(Iterator[T5]) Iterator[T6],
	f7 func(Iterator[T6]) Iterator[T7],
	f8 func(Iterator[T7]) Iterator[T8],
) Iterator[T8] {
	return f8(f7(f6(f5(f4(f3(f2(f1(in))))))))
}

// Pipe9 nine step pipe operator.
func Pipe9[T any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](
	in Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
	f2 func(Iterator[T1]) Iterator[T2],
	f3 func(Iterator[T2]) Iterator[T3],
	f4 func(Iterator[T3]) Iterator[T4],
	f5 func(Iterator[T4]) Iterator[T5],
	f6 func(Iterator[T5]) Iterator[T6],
	f7 func(Iterator[T6]) Iterator[T7],
	f8 func(Iterator[T7]) Iterator[T8],
	f9 func(Iterator[T8]) Iterator[T9],
) Iterator[T9] {
	return f9(f8(f7(f6(f5(f4(f3(f2(f1(in)))))))))
}

// Pipe10 ten step pipe operator.
func Pipe10[T any, T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](
	in Iterator[T],
	f1 func(Iterator[T]) Iterator[T1],
	f2 func(Iterator[T1]) Iterator[T2],
	f3 func(Iterator[T2]) Iterator[T3],
	f4 func(Iterator[T3]) Iterator[T4],
	f5 func(Iterator[T4]) Iterator[T5],
	f6 func(Iterator[T5]) Iterator[T6],
	f7 func(Iterator[T6]) Iterator[T7],
	f8 func(Iterator[T7]) Iterator[T8],
	f9 func(Iterator[T8]) Iterator[T9],
	f10 func(Iterator[T9]) Iterator[T10],
) Iterator[T10] {
	return f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(in))))))))))
}
