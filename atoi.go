package piscine

func Atoi(s string) int {

	runes := []rune(s)
	LenRune := 0
	result := 0
	for i := range runes {
		LenRune = i
	}
	if LenRune == 0 {
		return 0
	}

	tens := 1
	for k := 0; k < LenRune; k++ {
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
