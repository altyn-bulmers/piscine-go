package piscine

func TrimAtoi(s string) int {
	var runes []rune
	length := 0

	for i, letter := range s {

		if letter >= '0' && letter <= '9' {
			runes = append(runes, letter)
			length = i + 1
		} else if letter == '-' && length == 0 {
			runes = append(runes, '-')
		} else {
			continue

		}
	}
	return atoi(runes)
}

func atoi(runes []rune) int {

	LenRune := 0
	result := 0
	for i := range runes {
		i++
		LenRune++
	}
	if LenRune == 0 {
		return 0
	}

	tens := 1
	for k := 0; k < LenRune-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}
		tens *= 10
	}

	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		numb := 0
		for j := '0'; j < runes[i]; j++ {
			numb++
		}
		result += numb * tens
		tens /= 10
	}
	if runes[0] == '-' {
		result *= -1
	}
	return result
}
