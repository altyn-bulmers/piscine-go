package piscine

func BasicAtoi2(s string) int {
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
		if 48 > Bstr[i] || Bstr[i] > 57 {
			return 0
			break
		} else {
			Cstr[i] = Bstr[LenRune-1-i]
			c = 0
			for v := Cstr[i]; v >= '0'; v-- {
				c++
				count = Pow(10, i) * (c - 1)
				//fmt.Println(i, v, Cstr[i], c, count)

			}
			fin += count

		}
	}
	return fin
}
