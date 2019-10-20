package piscine

func IterativeFactorial(nb int) int {
	
	result := 1
	for i := 1; i < nb+1; i++ {
		result *= i
	}
	if nb < 0 {
		return 0
	} else {
		return result
	}
}
