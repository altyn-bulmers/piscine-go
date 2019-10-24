package piscine

func AppendRange(min, max int) []int {

	var array []int

	for i := min - 1; i < max-1; i++ {
		array = append(array, i+1)
	}
	return array
}
