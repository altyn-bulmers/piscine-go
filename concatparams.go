package piscine

func ConcatParams(args []string) string {
	var result string

	for _, v := range args {
		result += v + "\n"
	}
	return result

}
