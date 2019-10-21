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

func IsPrime(nb int) bool {
	result := true
	if nb <= 1{
		result = false
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				result = false
			} 
			
		} 
	}	
	return result 
}