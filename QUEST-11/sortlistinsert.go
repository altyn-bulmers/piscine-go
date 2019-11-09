package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI{
	l = listPushBackNode(l, data_ref)
	l = ListSort(l)
	return l
}

func listPushBackNode(l *NodeI, data int) *NodeI {
	n := &NodeI{Data:data, Next: nil}

	if l == nil {
		return n
	}
	a := l
	for a.Next != nil {
		a = a.Next
	}
	a.Next = n
	return l
}

func ListSort(l *NodeI) *NodeI {
	Head := l
	if Head == nil {
		return nil
	}

	Head.Next = ListSort(Head.Next)
	
	
	if Head.Next != nil && Head.Data > Head.Next.Data {
		Head = move(Head)
	}
	return Head	
}

func move(l *NodeI) *NodeI {
	p := l
	n := l.Next
	ret := n

	for n != nil && l.Data > n.Data {
		p = n
		n = n.Next
	}

	p.Next = l
	l.Next = n
	return ret
}