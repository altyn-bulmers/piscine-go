package piscine

func FirstRune(s string) rune {
	var runes []rune

	index := 0

	for _, letter := range str {

		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			runes = append(runes, letter)

		} else {
			continue

		}
	}
	return rune(runes[index])

}
