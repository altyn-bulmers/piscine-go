package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	length := 0
	for i := range tab {
		length = i + 1
	}

	asciending := true
	descending := true

	for i := 1; i < length; i++ {
		if !(f(tab[i-1], tab[i]) >= 0) {
			descending = false
		}
	}

	for i := 1; i < length; i++ {
		if !(f(tab[i-1], tab[i]) <= 0) {
			asciending = false
		}
	}

	return asciending || descending
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
