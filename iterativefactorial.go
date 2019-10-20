package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 1; i < nb+1; i++ {
		result *= i
	}
	if result > 0 {
		return result
	} else {
		result = 0
		return result
	}
}
