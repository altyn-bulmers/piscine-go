package piscine

func Rot14(str string) string {
	runes := []rune(str)
	var alphabet [52]rune
	var alphabig [52]rune

	for i, j, k := 0, 'a', 26; j <= 'z'; i, j, k = i+1, j+1, k+1 {
		alphabet[i] = j
		alphabet[k] = j
	}

	for i, j, k := 0, 'A', 26; j <= 'Z'; i, j, k = i+1, j+1, k+1 {
		alphabig[i] = j
		alphabig[k] = j
	}

	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			for q := 0; q < 26; q++ {
				alphindex := q
				if r == alphabet[alphindex] {
					runes[i] = alphabet[alphindex+14]
				}
			}

		} else if r >= 'A' && r <= 'Z' {
			for q := 0; q < 26; q++ {
				alphindex := q
				if r == alphabig[alphindex] {
					runes[i] = alphabig[alphindex+14]
				}
			}
		}
	}

	newStr := ""
	for _, r := range runes {
		newStr += string(r)
	}

	return newStr

}
