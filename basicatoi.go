package piscine

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func BasicAtoi(s string) int {
	Bstr := []rune(s)
	Cstr := []rune(s)

	LenRune := 0
	for i := range Bstr {
		i++
		LenRune++
	}
	count := 0
	fin := 0
	c := 0
	for i := 0; i < LenRune; i++ {
		Cstr[i] = Bstr[LenRune-1-i]
		c = 0
		for v := Cstr[i]; v >= '0'; v-- {
			c++
			count = Pow(10, i) * (c - 1)

		}
		fin += count
	}
	return fin
}
