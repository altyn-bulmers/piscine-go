package piscine

func SplitWhiteSpaces(str string) []string {
	index := 0
	count := 0
	word := ""
	for i, v := range str {
		if v == ' ' && str[i+1] != ' ' {
			count++
		}
	}
	result := make([]string, count+1)
	for _, r := range str {
		if isSeparator(r) {
			if word != "" {
				result[index] = word
				index++
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	size := 0
	for z := range result {
		size++
		z++
	}
	if word != "" {
		result[size-1] = word
	}
	return result
}

func isSeparator(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
