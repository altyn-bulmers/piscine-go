package main

import "fmt"

func main() {
	str := "HelloHAHowHAareHAyou?"
	charset := "HA"


	length := 0
	sublength := 0
	for l := range str {
		length = l + 1
	}
	for s := range charset {
		sublength = s + 1
	}

	for i:=0; i<sublength; i++{
		str += " "
	}

	len := 0
	prev := false
	for i:=0; i<length; i++{
		if str[i:i+sublength] == charset && !prev {
			len++
			prev = true
		} else {
			prev = false
		}
	}
	len++
	fmt.Println(len)
	array := make([]string, len)
	arrindex := 0
	word := ""
	pos := 0

	for i:=0; i<length+1; i++ {
		if str[i:i+sublength] == charset {
			word = str[pos:i]
			array[arrindex] = word
			arrindex++
			pos = i + sublength
		}
	}
	word = str[pos:length]
	array[arrindex] = word
//	if word != "" {
//		array[len-1] = word
//	}
	fmt.Println(array)
}