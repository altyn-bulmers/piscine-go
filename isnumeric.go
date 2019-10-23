package piscine

func checknum(a rune) bool {
	if a >= '0' && a <= '9' {
		return true
	}
	return false
}

func IsNumeric(str string) bool {
	runes := []rune(str)

	for i := range runes {
		if checknum(runes[i]) == false {
			return false
		}
	}
	return true
}
