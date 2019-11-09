package main

import "github.com/01-edu/z01"

//import "fmt"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {

	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	IntoRune(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	IntoRune(points.y)
	z01.PrintRune('\n')
	//	fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}

func check(r int) {

	c := '0'
	if r == 0 {
		z01.PrintRune(c)
		return
	}
	for i := 1; i <= r%10; i++ {
		c++
	}
	for i := -1; i >= r%10; i-- {
		c++
	}
	if r/10 != 0 {
		check(r / 10)
	}
	z01.PrintRune(c)
	return

}

func IntoRune(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	check(n)
}

/*
type student struct {
	name string
	age int
}

func changeName (pointer *student, nameChosen string) {

	pointer.name = nameChosen
}



	chris := student{}

	chris.name = "chris"
	chris.age = 30

	changeName(&chris, "lee")

	ptr := point{}

*/
