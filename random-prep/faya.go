package main

import "fmt"

func main(){

	first := []string{"faya"}
	second := []string{"fffgvyadgfdgr"}

	secondStr := ""
	firstStr := ""

	for _, s := range second	{
	//	fmt.Println(s)
		secondStr = secondStr + s
	}
	for _, s := range first	{
	//	fmt.Println(s)
		firstStr = firstStr + s
	}

//	fmt.Println(first, second)
//	fmt.Println(firstStr, secondStr)
	result := ""
	index := 0
	for _, s := range firstStr {
		
		for j := index; j< len( secondStr); j++ {
			if byte(s) == secondStr[j] {
				result = result + string(s)
				index = j + 1
				break
			}
		}
	}
	fmt.Println(result)

}