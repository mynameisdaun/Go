package main

import (
	"fmt"
)


func main() {
	
	//many ways to copy array
	slice1 := []int{1,2,3,4,5}
	slice2 := make([]int, len(slice1))

	for i,v := range slice1{
		slice2[i]=v
	}

	slice3 := append([]int{},slice1...)
	fmt.Println(slice3)

	slice4 := make([]int, 4)
	cnt1 := copy(slice4,slice1)
	fmt.Println(cnt1)
	fmt.Println(slice4)
	fmt.Println("I want sleep!")



	
}