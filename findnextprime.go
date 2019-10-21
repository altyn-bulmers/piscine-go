package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		n := 0
		for i := nb + 1; i < nb*2; i++ {
			n = i
			if IsPrime(i) == true {
				return n
				break
			}
		}
		return n
	}
}
