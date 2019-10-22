package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb%2 == 0 {
		nb += 1
	}
	for i := nb; ; i += 2 {
		if IsPrime(i) {
			return i
		}
	}
}
