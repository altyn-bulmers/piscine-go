package main

import "fmt"

func main() {
	str := "HelloHAHowHAareHAyou?"
	charset := "HA"

	length := 0
	sublength := 0

	for i := range str {
		length = i + 1
	}

	for i := range charset {
		sublength = i + 1
	}

	for i := 0; i < sublength; i++ {
		str += " "
	}
	prev := false
	len := 0
	for i := 0; i < length; i++ {
		if (str[i:i+sublength] == charset) && !prev {
			prev = true
			len++
		} else {
			prev = false
		}
	}
	len++
	fmt.Println(len)
	arr := make([]string, len)
	word := ""
	arindex := 0

	for i := 0; i < length; i++ {
		if str[i:i+sublength] == charset {
			l := 0
			for i := range word {
				l = i + 1
			}
			if l == 0 {
				continue
			}
			arr[arindex] = word
			arindex++
			word = ""
			i = i + sublength - 1
			continue
		}
		word += string(str[i])
	}

	l := 0
	for i := range word {
		l = i + 1
	}
	if l != 0 {
		arr[arindex] = word
	}
	fmt.Println( arr)
}
/*
	str = str + charset
	len := 0
	sepLen := 0

	for i := range str {
		len = i + 1
	}

	for i := range charset {
		sepLen = i + 1
	}

	
	
	
	wCount := 0
	for i := 0; i < len-sepLen+1; i++ {
		tmp := ""
		for j := 0; j < sepLen; j++ {
			tmp = tmp + string(str[i+j])
		}
		if tmp == charset {
			wCount++
		}
	}
	fmt.Println(wCount)
	res := make([]string, wCount)
	wCount = 0
	pos := 0
	for i := 0; i < len-sepLen+1; i++ {
		tmp := ""
		word := ""
		pl := 0
		for j := 0; j < sepLen; j++ {
			tmp = tmp + string(str[i+j])
			pl = j
		}
		if tmp == charset {
			word = str[pos:i]
			res[wCount] = word
			wCount++
			pos = i + pl + 1
		}
	}
	fmt.Println(res)
}
*/