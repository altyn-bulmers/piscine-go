package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := AiBase(nbr, baseFrom)
	return PNBase(n, baseTo)
}

func AiBase(s string, base string) int {

	len_s := 0
	for i := range s {
		len_s = i + 1
	}

	len_base := 0
	for i := range base {
		len_base = i + 1
	}

	if len_base < 2 {
		return 0
	}

	for i := 0; i < len_base-1; i++ {
		for j := i + 1; j < len_base; j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}

	count := 0
	var nums [40]int
	for i := range s {
		for j := range base {
			if s[i] == base[j] {
				nums[count] = j
				count++
			}
		}
	}

	res := 0
	for i := 0; i < len_s; i++ {
		res += nums[i] * pow(len_base, len_s-1-i)
	}
	return res

}

func pow(a int, b int) int {
	if b < 0 {
		return 0
	}
	ans := 1
	for b != 0 {
		ans *= a
		b--
	}
	return ans
}

func PNBase(nbr int, base string) string {
	length := 0
	for i := range base {
		length = i + 1
	}

	if length < 2 {
		return ""
	}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if base[i] == base[j] {
				return ""
			}
		}
	}

	var nums [70]int
	var runes [70]rune
	count := 0
	length_num := 0
	temp := nbr
	for temp != 0 {
		length_num++
		temp = temp / length
	}

	for nbr != 0 {
		rem := nbr % length
		nums[length_num-1-count] = rem
		count++
		nbr = nbr / length
	}

	base_runes := []rune(base)
	c := 0
	for i := 0; i < length_num; i++ {
		for j := range base_runes {
			if j == nums[i] {
				runes[c] = base_runes[j]
				c++
			}
		}
	}

	return string(runes[:length_num])
}
