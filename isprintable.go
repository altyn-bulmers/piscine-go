package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)

	for i := range runes {
		if check(runes[i]) == false {
			return false
		}
	}
	return true
}
