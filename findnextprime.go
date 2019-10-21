package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	}
	return FindNextPrime(nb + 1)
}
