package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	length := 0
	for i := range s {
		length = i + 1
	}
	return runes[length-1]
}
