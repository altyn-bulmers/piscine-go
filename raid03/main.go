package main

import (
	"bufio"
	"fmt"
	"os"
	//    	"sort"
)

func counter(output []rune) (x, y int) {
	countX := 0
	countY := 0
	flag := true
	for _, s := range output {
		if s == '\n' {
			countY++
			flag = false
		} else {
			if flag == true {
				countX++
			}
		}
	}
	return countX, countY
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var output []rune
	inputStr := ""
	str := []string{}
	for reader.Scan() {
		inputStr = inputStr+reader.Text() + "\n"
	}
	
	//if size == 0 --> not a raid func
	for _, a := range inputStr {
		output = append(output, a)
	}
	x, y := counter(output)
	//determineInput(output)
	strX := printNbr(x)
	strY := printNbr(y)

	if raid1a(x, y) == inputStr {
		str = append(str, "[raid1a]" + " "  + "["+strX+"]"+" "+"["+strY+"]" + " ")
	}
	if raid1b(x, y) == inputStr {
		str = append(str, "[raid1b]" + " " + "["+strX+"]"+" "+"["+strY+"]" + " ")
	}
	if raid1c(x, y) == inputStr {
		str = append(str, "[raid1c]" + " " + "["+strX+"]"+" "+"["+strY+"]" + " ")
	}
	if raid1d(x, y) == inputStr {
		str = append(str, "[raid1d]" + " " + "["+strX+"]"+" "+"["+strY+"]" + " ") 
	}
	if raid1e(x, y) == inputStr {
		str = append(str, "[raid1e]" + " " + "["+strX+"]"+" "+"["+strY+"]" + " ")
	}
	if len(str) == 0 {
		fmt.Println("Not a Raid function")
		return
	}
	for i := 0; i <= len(str)-2; i++ {
		fmt.Print(str[i] + "||" + " ")
	}
	fmt.Println(str[len(str)-1])
}

func raid1a(x, y int) string {
	result := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (j == 1 && i == 1) || (j == 1 && i == y) {
				result = result + "o"
			} else if (j == x && i == y) || (j == x && i == 1) {
				result = result + "o"
			} else if j == 1 || j == x {
				result = result + "|"
			} else if i == 1 || i == y {
				result = result + "-"
			} else {
				result = result + " "
			}
			if j == x {
				result = result + "\n"

			}

		}

	}

	return result
}

func raid1b(x, y int) string {
	result := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (j == 1 && i == 1) || (j == x && i == y) {
				result = result + "/"
			} else if (j == 1 && i == y) || (j == x && i == 1) {
				result = result + "\\"
			} else if j == 1 || j == x {
				result = result + "*"
			} else if i == 1 || i == y {
				result = result + "*"
			} else {
				result = result + " "
			}
			if j == x {
				result = result + "\n"

			}

		}

	}

	return result
}

func raid1c(x, y int) string {
	result := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (j == 1 && i == 1) || (j == x && i == 1) {
				result = result + "A"
			} else if (j == x && i == y) || (j == 1 && i == y) {
				result = result + "C"
			} else if j == 1 || j == x {
				result = result + "B"
			} else if i == 1 || i == y {
				result = result + "B"
			} else {
				result = result + " "
			}
			if j == x {
				result = result + "\n"

			}

		}

	}

	return result
}

func raid1d(x, y int) string {
	result := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (j == 1 && i == 1) || (j == 1 && i == y) {
				result = result + "A"
			} else if (j == x && i == y) || (j == x && i == 1) {
				result = result + "C"
			} else if j == 1 || j == x {
				result = result + "B"
			} else if i == 1 || i == y {
				result = result + "B"
			} else {
				result = result + " "
			}
			if j == x {
				result = result + "\n"

			}

		}

	}

	return result
}

func raid1e(x, y int) string {
	result := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (j == 1 && i == 1) || (j == x && i == y) {
				result = result + "A"
			} else if (j == 1 && i == y) || (j == x && i == 1) {
				result = result + "C"
			} else if j == 1 || j == x {
				result = result + "B"
			} else if i == 1 || i == y {
				result = result + "B"
			} else {
				result = result + " "
			}
			if j == x {
				result = result + "\n"

			}

		}

	}

	return result
}

func printNbr(n int) string {
	res := ""

	if n/10 != 0 {
		res = res + printNbr(n/10)
	}
	mod := '0'
	for i := 0; i < (n % 10); i++ {
		mod++

	}
	if mod == '0' {
	res = res + "0"
} else if mod == '1' {
	res = res + "1"
}else if mod == '2' {
	res = res + "2"
}else if mod == '3' {
	res = res + "3"
}else if mod == '4' {
	res = res + "4"
}else if mod == '5' {
	res = res + "5"
}else if mod == '6' {
	res = res + "6"
}else if mod == '7' {
	res = res + "7"
}else if mod == '8' {
	res = res + "8"
}else if mod == '9' {
	res = res + "9"
}

	return res

}
