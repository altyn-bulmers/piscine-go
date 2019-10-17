package piscine

func StrLen(str string) {
	Astr := []rune(str)
	count := 0
	for i := range Astr {
		i++
		count++
	}
	return count

}
