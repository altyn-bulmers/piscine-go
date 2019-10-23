package piscine

func checkup(a rune) bool {
	if a >= 'A' && a <= 'Z' {
		return true
	}
	return false
}

func IsUpper(str string) bool {
	runes := []rune(str)

	for i := range runes {
		if checkup(runes[i]) == false {
			return false
		}
	}
	return true
}
