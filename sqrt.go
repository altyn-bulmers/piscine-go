package main

import "fmt"



func main(){
	nb := 15
	if nb <0 || nb*10%10 != 0{
		fmt.Println("zero")
	} else {
		for i:=0; i<101; i++{
			if nb == i*i{
				nb = i
				break
			} else if i*i > nb {
				nb = 0
				break
			}
		}
		fmt.Println(nb)
	}

	
}





/*
func Sqrt(nb int) int{

	result := 1
	nb := 16

	if result*10%10 != 0{
		return 0
	} else {
		for i:=0; i<20; i++ {
			result -= (result*result - nb) / (2*result)
			
		}
		return result
	}



}
*/

