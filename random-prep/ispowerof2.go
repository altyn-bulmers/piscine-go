package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := []string{"2"}

	length := 0
	for l := range str {
		length = l + 1
	}

	if length != 1 {
		return
	}
	strNumber := str[0]

	number, err := strconv.Atoi(strNumber)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	x := 1
	flag := false
	//if number == 2 {
	//	flag = true
	//}
	for i := 1 ; i <= 15; i++ {
		x = x*2
		fmt.Println(x)
		if x == number {
			flag = true
		//	fmt.Println("power: ", i)
			break 
		}
	}
	if flag == true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
