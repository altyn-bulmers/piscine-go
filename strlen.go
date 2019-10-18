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
