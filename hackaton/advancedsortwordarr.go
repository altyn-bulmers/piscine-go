package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	temp := ""
	count := 0
	for c := range array {
		count = c + 1
	}

	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if f(array[i], array[j]) > 0 {
				temp = array[i]
				array[i] = array[j]
				array[j] = temp

			}
		}

	}

}
