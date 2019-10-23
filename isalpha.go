package piscine

func check(a rune) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func IsAlpha(str string) bool {
	runes := []rune(str)

	for i := range runes {
		if check(runes[i]) == false {
			return false
		}
	}
	return true
}
