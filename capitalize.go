package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	length := 0
	for j := range runes {
		j++
		length++
	}

	i := 1
	for i < length+1 {
		if (runes[i-1] == ' ' || runes[i-1] == '+' || i == 0) && (runes[i] >= 'a' && runes[i] <= 'z') {
			runes[i] = runes[i] - 32
		}
		i++
	}
	return string(runes)
}
