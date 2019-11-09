package main

import "fmt"

func main() {
	str:="HelloHAHowHAareHAyou?"
	charset := "HA"
	fmt.Println(Split(str, charset))
}

func Split(str, charset string) []string {
	str = str + charset
	len := strLen(str)
	sublen := strLen(charset)
	wCount := 0
	for i:=0; i < len - sublen + 1; i++{
		temp := ""
		for j:=0; j< sublen; j++{
			temp = temp + string(str[i+j])
		}
		if temp == charset {
			wCount++
		}
	}

	res := make([]string, wCount)
	wCount = 0
	pos:= 0
	for i:=0; i < len - sublen + 1; i++{
		temp := ""
		word := ""
		pl := 0
		for j:=0; j<sublen;j++{
			temp = temp + string(str[i+j])
			pl = j
		}
		if temp == charset {
			word = str[pos:i]
			res[wCount] = word
			wCount++
			pos = i + pl + 1
		}
	} 
	return res
}

func strLen(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}