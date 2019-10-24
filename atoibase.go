package piscine

func AtoiBase(s string, base string) int {
	baseF := Length(base)
	result := []rune(base)
	pos := 0

	if constrains1(baseF, result) == true {
		for i := 0; i < Length(s); i++ {
			pos *= Length(base)
			pos += Index1(s, base)
		}
		return pos
	}
	return 0
}

func constrains1(baseF int, result []rune) bool {

	if baseF == 0 {
		return false
	}
	for a := 0; a < Length(string(result)); a++ {
		for x := a + 1; x < Length(string(result)); x++ {
			if baseF < 2 || result[a] == result[x] || result[x] == '+' || result[x] == '-' {
				return false
			}
		}
	}
	return true
}

func Index1(s string, toFind string) int {
	arrayStr := []rune(s)
	array := []rune(toFind)
	for j := 0; j < Length(s); j++ {
		for x := 0; x < Length(toFind); x++ {

			if arrayStr[j] == array[x] { //[0:3]
				return x
			}
		}
	}

	return -1
}

func Length(a interface{}) int {
	count := 0
	switch a := a.(type) {
	case string:
		for range a {
			count++
		}
		break
	case []string:
		for range a {
			count++
		}
	case []int:
		for range a {
			count++
		}
	case []rune:
		for range a {
			count++
		}
	}
	return count
}
