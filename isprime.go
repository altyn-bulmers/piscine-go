package piscine

func IsPrime(nb int) bool {
	result := true
	if nb <= 1 {
		result = false
	} else {
		for i := 2; i <= Sqrt(nb); i++ {
			if nb%i == 0 {
				result = false
			}

		}
	}
	return result
}
