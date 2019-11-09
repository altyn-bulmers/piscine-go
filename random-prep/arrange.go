package main

import (
//	"github.com/01-edu/z01"
	"fmt"
)

func main(){
	
	fmt.Println(arrange(4,0))



}

// arrange from second number to the first
/*
3.0.
0 1 2 3 
0 3
3 2 1 0
2 -3 
-3 -2 -1 0 1 2    */
func arrange(first int, second int) []int {

	length := 0
	if second > first && (second < 0 && first < 0) {
		length = (first - second) * -1
	} else if second <0 {
		length = (second - first ) * -1
	} else if second < first {
		length = first - second
	} else {
		length = second - first
	}

	array := make([]int, length+1)
	
	for i := second; i <= first; i++{
		array[i-second] = i
	}

	for i := second; i >= first; i--{
		array[second-i] = i
	}

	

	return array
}