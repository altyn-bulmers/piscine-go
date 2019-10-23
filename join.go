package piscine

func Join(strs []string, sep string) string {

	var result string
	for i := range strs {
		result += strs[i]
		result += sep
	}

	return result

}
