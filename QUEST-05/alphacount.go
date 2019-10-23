package piscine

func AlphaCount(str string) int {
	runes := []rune(str)
	count := 0
	for _, letter := range runes {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			count++
		}
	}
	return count
}
