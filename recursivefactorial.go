package piscine

func RecursiveFactorial(nb int) int {
	if nb > 1 && nb < 2147483647 {
		return RecursiveFactorial(nb-1) * nb
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		return 0
	}
}
