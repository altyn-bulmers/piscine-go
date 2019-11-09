package piscine

func Unmatch(arr []int) int {

	var qount int

	for _, el := range arr {
		qount = 0
		for _, v := range arr {
			if v == el {
				qount++
			}
		}
		if qount%2 != 0 {
			return el
		}
	}
	return -1
}
