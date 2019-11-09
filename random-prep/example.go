package main 

import "fmt"

type node struct {
	data int
	next *node
}

func insert(head *node, data int) *node {
	n := &node {data:data}

	if head == nil {
		return n
	} else {
		n.next = head
		return n
	}
}

func PrintList(head *node) {
	for head != nil {
		fmt.Print(head.data, " -> ")
		head = head.next
	}
	fmt.Println(nil)
}

func main() {
	var link *node

	link = insert(link, 1)
	link = insert(link, 2)
	link = insert(link, 3)

	PrintList(link)
}