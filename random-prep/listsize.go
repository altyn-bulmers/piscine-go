package main

import "fmt"

func main(){


}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize (l *List) int {
	
	size := 0
	for l.Head != nil {
		size++
		l.Head = l.Head.Next
	}
	return size
}