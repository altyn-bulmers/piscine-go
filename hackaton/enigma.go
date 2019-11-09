package piscine

/*
func Enigma(a ***int, b *int, c *******int, d ****int) {
	y := *******c

	*******c = ***a
	t := ****d
	****d = y

	g := *b
	*b = t
	***a = g

}
*/

func Enigma(a ***int, b *int, c *******int, d ****int) {
	*******c, ****d, *b, ***a = ***a, *******c, ****d, *b
}
