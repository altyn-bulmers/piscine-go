package piscine

func BasicAtoi3(s string) int {
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
			//fmt.Println(i, v, Cstr[i], c, count)
			//fmt.Println(Cstr[LenRune-1])
		}
		fin += count

		if 48 > Bstr[i] || Bstr[i] > 57 {
			return 0
			break
		}

	}
	return fin
}
func PlusMinus(s string) int {
	Bstr := []rune(s)
	Cstr := []rune(s)

	LenRune := 0
	for i := range Bstr {
		i++
		LenRune++
	}
	Kstr := Bstr[1:LenRune]
	count := 0
	fin := 0
	c := 0
	for i := 0; i < LenRune-1; i++ {
		Cstr[i] = Kstr[LenRune-2-i]

		c = 0
		for v := Cstr[i]; v >= '0'; v-- {
			c++
			count = Pow(10, i) * (c - 1)
			//fmt.Println(i, v, Cstr[i], c, count)
			//fmt.Println(Cstr[LenRune-1])
		}
		fin += count
		if 48 > Kstr[i] || Kstr[i] > 57 {
			return 0
			break
		}

	}
	return fin
}
func Atoi(s string) int {
	result := 0

	Mstr := []rune(m)

	if Mstr[0] == 43 { // "+"
		result = PlusMinus(m)
	} else if Mstr[0] == 45 { // "-"
		result = 0 - PlusMinus(m)
	} else {
		result = BasicAtoi3(m)
	}
	return result
}
