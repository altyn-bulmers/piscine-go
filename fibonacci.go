package piscine

func Fibonacci(index int) int {

	if index >= 3 {
		return Fibonacci(index-1) + Fibonacci(index-2)

	} else if index == 0 {
		return 0
	} else if index < 0 {
		return -1
	}
	return 1
}
