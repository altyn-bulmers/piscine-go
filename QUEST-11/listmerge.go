package piscine

/*
func ListMerge(l1 *List, l2 *List) {
	a := l1.Head
	b := l2.Head

	for a == nil && l2 != nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	}

	for a != nil {
		if a.Next == nil {
			a.Next = b
			return
		}
		a = a.Next
	}
}
*/

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	} else if l2.Head == nil {
		return
	}

	a := l1.Head
	for a.Next != nil {
		a = a.Next
	}
	a.Next = l2.Head
	l1.Tail = l2.Tail

}
