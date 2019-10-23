package piscine

func checklow(a rune) bool {
	if a >= 'a' && a <= 'z' {
		return true
	}
	return false
}

func IsLower(str string) bool {
	runes := []rune(str)

	for i := range runes {
		if checklow(runes[i]) == false {
			return false
		}
	}
	return true
}
