package it

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type Ordered interface {
	Integer | Float | ~string
}

type Yield[T any] struct {
	Value T
	Done  bool
}

type Iterator[T any] interface {
	Next() Yield[T]
}

type Generator[T any] interface {
	Iterator() Iterator[T]
}

type nextIterator[T any] struct {
	next func() Yield[T]
}

func (it *nextIterator[T]) Next() Yield[T] {
	return it.next()
}

type emptyIterator[T any] struct{}

func (it *emptyIterator[T]) Next() Yield[T] {
	return Yield[T]{Done: true}
}

type PrimeFactor[T Number] struct {
	Prime T
	Power int
}
