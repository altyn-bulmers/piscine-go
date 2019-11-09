package piscine

func SortWordArr(array []string) {
	count := 0
	for c := range array {
		count = c + 1
	}

	temp := 0
	for i := 0; i < count-1; i++ {
		temp = i
		for j := i + 1; j < count; j++ {
			if array[j] < array[temp] {
				temp = j
			}
		}
		array[i], array[temp] = array[temp], array[i]
	}

}
