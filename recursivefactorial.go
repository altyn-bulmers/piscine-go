package piscine

func RecursiveFactorial(nb int) int {
	if nb > 1 {
		return RecursiveFactorial(nb -1) * nb
	} else {
		return 1
	}
}
