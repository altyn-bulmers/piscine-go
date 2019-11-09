package main

import (
	piscine ".."
	"fmt"
)

func main() {
	var donnie piscine.Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = 1

	fmt.Println(donnie)
}

