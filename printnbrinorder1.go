package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder1(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var cnt [10]int
	for n != 0 {
		cnt[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {
		for cnt[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			cnt[i]--
		}
	}
}
