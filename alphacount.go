package piscine

func AlphaCount(str string) int {
	runes := []rune(str)
	count := 0
	for i, letter := range runes {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			count++
		}
	}
	return count
}
