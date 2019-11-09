package main 

import "fmt"

func main (){
	fmt.Println(ItoaBase(-125, 4))


}

func ItoaBase(value, base int) string {
	origin := "0123456789ABCDEF"
	str := origin
	str = str[:base] 
	runes := []rune(str)
	valid := true
	result := ""
	if base < 2 {
		valid = false
	}
	if !valid {
		return "NV"
	}

	if value == 0 {
		return "0"
	} else {
		if value < 0 {
			result = result + "-"
		}
		result += BaseRecursion(value, runes, base)
	}
	return result
}

func BaseRecursion(nbr int, runes []rune, base int) string {
	result := ""
	if nbr/base != 0 {
		result = BaseRecursion (nbr/base, runes, base) 
	}
	mod := 0
	mod = nbr % base 
	if mod <0 {
		mod = -mod
	}
	result = result + string(runes[mod]) 
	
	return result
}

