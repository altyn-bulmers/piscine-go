package piscine

func Join(strs []string, sep string) string {
	var result string
	var length int

	for j := range strs {
		j++
		length++
	}

	for i := range strs {
		result += strs[i]
		if i != length-1 {
			result += sep
		}
	}

	return result

}
