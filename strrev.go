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
	Bstr := []rune(str)
	c := StrLen(str)
	d := []rune()
	for i := c - 1; i >= 0; i-- {
		d := Bstr[i]
	}
	return string(d)
}
