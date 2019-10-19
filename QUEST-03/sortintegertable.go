package piscine

func SortIntegerTable(table []int) {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			tmp := table[i-1]
			table[i-1] = table[i]
			table[i] = tmp
			i = 1
		} else {
			i++

		}
	}
}
