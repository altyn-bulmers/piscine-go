package main

func IterativeFactorial(nb int) int {
	result := 1
	for i := 1; i < nb+1; i++ {
		result *= i
	}
	if result > 0 {
		return result
	} else {
		return 0
	}
}

func main() {
	arg := 4
	IterativeFactorial(arg)
}
