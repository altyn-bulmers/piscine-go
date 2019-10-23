package piscine

func Index(s string, toFind string) int {
	length := 0
	sublength := 0

	for i := range s {
		length = i + 1
	}

	for i := range toFind {
		sublength = i + 1
	}

	t := 0

	for i := 0; i < length; i++ {
		j := 0
		t = i
		for j < sublength {
			if t < length && s[t] == toFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == sublength {
			return i
		}
	}
	return -1
}
