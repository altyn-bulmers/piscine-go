package piscine

func check(a rune) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	runes := []rune(s)

	first := true

	for i := range runes {
		if check(runes[i]) == true && first {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] -= 32
			}
			first = false
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		} else if check(runes[i]) == false {
			first = true
		}
	}
	return string(runes)
}
