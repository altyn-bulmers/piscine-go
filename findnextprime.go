package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		n := 0
		for i := nb + 1; i < 2147483647; i++ {
			n = i
			if IsPrime(i) == true {
				return n
				break
			}
		}
		return n
	}
}

func IsPrime(n int) bool {
	result := true
	if n <= 1{
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
