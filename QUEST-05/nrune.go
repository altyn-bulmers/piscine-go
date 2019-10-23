package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	runes := []rune(s)
	length := 0
	for i := range s {
		length = i + 1
	}
	if length == 0 || n > length {
		return 0
	}

	return runes[n-1]
}
