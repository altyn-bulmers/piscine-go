package main

import "fmt"

func main() {
	fmt.Println(ItoaBase(-125, 10))
}


func ItoaBase(value, base int) string {
	origin := "0123456789ABCDEF"
	str := origin[:base]
	runes := []rune(str)
	valid := true
	result := ""
	if base < 2 {
		valid = false
	}
	if !valid {
		return "NV"
	} else {
		if value == 0 {
			return "0"
		} else {
			if value <0 {
				result = "-"
			}
			result += BaseRecursion(value, runes, base)
		}
	}
	return result 

}

func BaseRecursion(n int, runes []rune, len int) string {
	result := ""
	if n/len != 0 {
		result = BaseRecursion(n/len, runes, len)
	}
	mod := 0
	mod = n % len
	if mod < 0 {
		mod = -mod
	}
	result = result + string(runes[mod]) 
	return result
}