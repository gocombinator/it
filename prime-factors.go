package it

func primeFactorsAux[N Unsigned](value N) chan PrimeFactor[N] {
	var out = make(chan PrimeFactor[N])
	go func() {
		for prime := range primesAux[N]() {
			var exponent int = 0
			for {
				if (value % prime) != 0 {
					break
				}
				value /= prime
				exponent++
			}
			if exponent > 0 {
				out <- PrimeFactor[N]{prime, exponent}
			}
			if value <= 1 {
				break
			}
		}
		close(out)
	}()
	return out
}

// PrimeFactors yields prime factors of provided number.
func PrimeFactors[T Unsigned](n T) Iterator[PrimeFactor[T]] {
	return OfChan(primeFactorsAux(n))
}
