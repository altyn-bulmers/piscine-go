package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		n := nb + 1
		for IsPrime(n) == false {
			n++
		}
		return n
	}
}

func IsPrime(n int) bool {
	result := true
	if n <= 1 {
		result = false
	} else {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				result = false
			}

		}
	}
	return result
}