package it

func primesAux[N Unsigned]() <-chan N {
	var out = make(chan N)
	go func() {
		var m = make(map[N]N)
		var a N = 2
		for {
			if b, ok := m[a]; ok {
				delete(m, a)
				var c = a + b
				for {
					if _, ok := m[c]; !ok {
						break
					}
					c += b
				}
				m[c] = b
			} else {
				m[a*a] = a
				out <- a
			}
			a++
		}
	}()
	return out
}

// Primes yields prime numbers.
func Primes[T Unsigned]() Iterator[T] {
	return OfChan(primesAux[T]())
}
