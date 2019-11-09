package main

import "fmt"

func main() {
	s := []int{5,4,3,2,1,0}
	SortIntegerTable(s)
	fmt.Println(s)
}

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i<n; i++{
		for j:=i;j<n;j++{
			if table[i]>table[j]{
				Swap(&table[i],&table[j])
			}
		}
	}
}

func Swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
