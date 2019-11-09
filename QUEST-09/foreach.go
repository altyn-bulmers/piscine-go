package piscine

func ForEach(f func(int), arr []int) {

	for i := range arr {
		f(arr[i])
	}
	return

}
