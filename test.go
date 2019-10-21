package main

import (
	"fmt"
)

func main(){
	fmt.Println(Sqrt(16))
}

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	} else {
		for i := 0; i < nb/2; i++ {
			if i*i == nb  {
				return i
			} 
		}
	}
	return 0
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i <= Sqrt(nb); i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
