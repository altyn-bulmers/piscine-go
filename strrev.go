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
	RevRune := []rune(s)
	for i := LenRune - 1; i >= 0; i-- {
		RevRune = Bstr[i]
	}
	return string(RevRune)
}
