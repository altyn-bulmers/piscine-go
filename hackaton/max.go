package piscine

func Max(arr []int) int {
	lenght := 0
	for l := range arr {
		lenght = l + 1
	}
	i := 1
	for i < lenght {
		if arr[i-1] > arr[i] {
			temp := i
			arr[i] = arr[i-1]
			arr[i-1] = arr[temp]
			i = 1
		} else {
			i++
		}
	}

	return arr[lenght-1]
}
