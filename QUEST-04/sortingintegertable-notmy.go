package piscine

func partition(table []int, l int, r int) int {
	var pivot int = table[l + (r-l)/2]
	for l <= r {
		for table[l] < pivot {
			l++
		}
		for table[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		Swap(&table[l], &table[r])
		l++
		r--
	}
	return r
}

func quickSort(table []int, l int, r int) {
	if l < r {
		pi := partition(table, l, r)
		quickSort(table, l, pi-1)
		quickSort(table, pi+1, r)
	}
}

func SortIntegerTable(table []int) {
	quickSort(table, 0, len(table)-1)
}