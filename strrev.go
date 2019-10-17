package piscine

func StrLen(str string) int {
	Astr := []rune(str)
	count := 0
	for i := range Astr {
		i++
		count++
	}
	return count

}
func StrRev(s string) string {
	LenRune := StrLen(s)
	Bstr := []rune(s)
	var reverse string
	for i := LenRune - 1; i >= 0; i-- {
		reverse += string(Bstr[i])
	}
	return reverse
}
