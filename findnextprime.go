package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		if IsPrime(nb) == true || nb > 2147483647 || nb < -2147483648 {
			return nb
		}
	}
	return FindNextPrime(nb + 1)
}
