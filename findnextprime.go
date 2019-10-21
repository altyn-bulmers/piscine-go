package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		n := nb + 1
		for IsPrime(n) == true {
			break
		}
		return n
	}
}
