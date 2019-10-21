package piscine

// FindNextPrime function returns the first
// prime number that is equal or superior
// to the int passed as parameter.
// The function is optimized to avoid time-outs.
func FindNextPrime(nb int) int {
	var nextprime int

	if IsPrime(nb) == true {
		nextprime = nb
	} else {
		for i := nb + 1; i < nb*2; i++ {
			if IsPrime(i) == true {
				nextprime = i
				return nextprime
			}
		}
	}
	return nextprime
}

// IsPrime function returns true if the int
// passed as parameter is a prime number.
// Otherwise it returns false.
// The function is optimized to avoid time-outs.
