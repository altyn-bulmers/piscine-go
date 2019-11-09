package main

import "fmt"

func main() {
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
}

func AtoiBase(s string, base string) int {
	n := strLen(base)
	count := strLen(s)
	if n < 2 {
		return 0
	}
	for i := 0; i < n; i++ {
		if base[i] == '-' || base[i] == '+' {
			return 0
		}
		for j := i +1 ; j< n; j++ {
			if base[j] == base[i]{
				return 0
			}
		}
	}

	res := 0
	bn := 1
	numeric := 0

	for i := count -1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if base[j] == s[i] {
				numeric = j
			}
		} 
		res = res + bn * numeric
		bn = bn * n
	}
	return res
}

func strLen(s string) int {
	length := 0
	for i := range s {
		length = i + 1
	}
	return length
}