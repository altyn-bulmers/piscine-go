package piscine

func Compact(ptr *[]string) int {
	length := 0
	for _, s := range *ptr {
		if s != "" {
			length++
		}
	}
	count := 0
	array := make([]string, length)
	for _, s := range *ptr {
		if s != "" {
			array[count] = s
			count++
		}
	}
	*ptr = array
	return length
}
