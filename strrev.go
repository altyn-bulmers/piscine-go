package piscine

func StrRev(s string) string {
	Bstr := []rune(s)
	LenRune := 0
	for i := range Bstr {
		i++
		LenRune++
	}
	var reverse string
	for i := LenRune - 1; i >= 0; i-- {
		reverse += string(Bstr[i])
	}
	return reverse
}
