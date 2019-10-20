package piscine

func IterativeFactorial(nb int) int {
	if x < 0 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		result := 1
		for i := 1; i < nb+1; i++ {
			result *= i
			if result > 9223372036854775807 {
				result = 0
				break		
			}
		return result
		}
	}
