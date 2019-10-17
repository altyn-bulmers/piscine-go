package piscine

func BasicAtoi3(m string) int {
	Bstr := []rune(m)
	Cstr := []rune(m)

	LenRune := 0
	for i := range Bstr {
		i++
		LenRune++
	}
	if LenRune == 0 {
		return 0
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
func PlusMinus(m string) int {
	Bstr := []rune(m)
	Cstr := []rune(m)

	LenRune := 0
	for i := range Bstr {
		i++
		LenRune++
	}
	if LenRune == 0 {
		return 0
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

	Mstr := []rune(s)

	if Mstr[0] == 43 { // "+"
		result = PlusMinus(s)
	} else if Mstr[0] == 45 { // "-"
		result = 0 - PlusMinus(s)
	} else {
		result = BasicAtoi3(s)
	}
	return result
}
