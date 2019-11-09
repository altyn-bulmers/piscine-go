package piscine

func ListClear(l *List) {
	if l.Head != nil || l.Tail != nil {
		l.Head = nil
		l.Tail = nil
	}

}
