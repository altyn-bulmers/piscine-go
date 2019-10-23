package piscine

func BasicJoin(strs []string) string {
	var result string
	for i := range strs {
		result += strs[i]
	}

	return result

}
