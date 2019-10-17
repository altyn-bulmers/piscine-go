package piscine

func Atoi(s string) int {
	runes := []rune(s)

	num := 0
	length := 0
	for i := range runes {
		length = i
	}
	if length == 0 {
		return 0
	}
	ten := 1
	for j := 0; j < length; j++ {
		if runes[j] == '+' || runes[j] == '-' {
			continue
		}
		ten *= 10
	}
	for i := range runes {
		if (runes[0] == '+' || runes[0] == '-') && i == 0 {
			continue
		}
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		temp := 0
		for k := '0'; k < runes[i]; k++ {
			temp++
		}
		num += temp * ten
		ten /= 10
	}
	if runes[0] == '-' {
		num *= -1
	}
	return num
}
