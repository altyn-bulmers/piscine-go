package piscine

func UltimateDivMod(a *int, b *int) {
	y := *a
	*a = *a / *b
	*b = y % *b

}
