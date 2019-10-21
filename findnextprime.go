package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb%2 == 0 {
		nb += 1
	}
	for i := nb; ; i += 2 {
		if Prime(i) {
			return i
		}
	}
}

func abs(nb int) int {
	if nb < 0 {
		return -nb
	}
	return nb
}

func SQ(nb int) int {
	if nb <= 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	var closest int = 0
	for i := 1; i <= nb/2; i++ {
		if abs(nb-(i*i)) < abs(nb-(closest*closest)) {
			closest = i
		} else {
			break
		}
	}
	return closest
}

func Prime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i <= SQ(nb); i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
