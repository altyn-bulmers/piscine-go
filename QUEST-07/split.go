package piscine

func Split(str, charset string) []string {
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
	return arr
}
