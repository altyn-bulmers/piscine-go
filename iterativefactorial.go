package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 1; i < nb+1; i++ {
		if result > 9223372036854775807 {
			result *= i
			break
			return 0
		} else {
			continue
		}
	}
	if nb < 0 {
		return 0
	} else {
		return result
	}
}
