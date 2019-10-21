package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i <= nb/3; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
