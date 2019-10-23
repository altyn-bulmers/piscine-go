package piscine

func checkprint(a rune) bool {
	if a >= 0 && a <= 31 {
		return true
	}
	return false
}

func IsPrintable(str string) bool {
	runes := []rune(str)

	for i := range runes {
		if checkprint(runes[i]) == false {
			return false
		}
	}
	return true
}
