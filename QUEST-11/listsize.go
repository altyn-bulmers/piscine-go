package piscine

func ListSize(l *List) int {
	size := 0

	for l.Head != nil {
		size++
		l.Head = l.Head.Next
	}

	return size
}
