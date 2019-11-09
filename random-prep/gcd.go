package main 

import (
	"fmt"
)

func main() {
	first := "42"
	second := "12"
	firstInt := atoi(first)
	secondInt := atoi(second)
	firstArr := primes(firstInt)
	secondArr := primes(secondInt)
	firstArr = makeUnique(firstArr)
	secondArr = makeUnique(secondArr)
	result := 1
	for _, f := range firstArr {
		for _, s := range secondArr{
			if s == f {
				result *= f
			}
		}
	}
	fmt.Println("first: ", firstArr)
	fmt.Println("second: ", secondArr)
	fmt.Println(result)
	

}

func makeUnique(slice []int) []int {
	new := []int{}
	flag := false
	for i, s := range slice {
		for j, r := range slice {
			if s == r && i != j {
				new = append(slice[:i], slice[i+1:]...)
				flag = true
			}
		}	
	}
	if flag == true {
		return new
	} else {
		return slice 
	}
	   
}

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb % 2 == 0 {
		return false
	}
	for i := 3; i <= nb/3; i +=2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func primes (n int) []int {
	var array []int
	for n%2 == 0 {
		array = append(array, 2)
		n /= 2
	}

	for i:=3; i*i <= n; i +=2 {
		if n%i == 0 {
			array = append(array, i)
		//	fmt.Print("n: ", n, " " )
			n /= i

		}
	}
	if isPrime(n) && n > 2 {
		array = append(array, n)
	}
	return array
}

func atoi(s string) int {
	runes := []rune(s)
	numb := 0
	for _, s := range runes{
		for i := '0'; i <= '9'; i++ {
			if s == i {
				numb = numb * 10 + int(i - '0')

			}
		} 
	}

	if runes[0] == '-' {
		numb *= -1
	}
	return numb










	
}