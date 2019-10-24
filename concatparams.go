package piscine

func ConcatParams(args []string) string {
	var result string
	size := 0
	for i := range args {
		i++
		size++
	}

	for i, v := range args {
		result += v
		if i != size-1 {
			result += "\n"
		}

	}
	return result
}
