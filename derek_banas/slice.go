package main

import "fmt"

func main() {
	mySlice := []int {2,3,5}
	mySlice2 := mySlice[1:3]
	fmt.Println(mySlice2)

	mySlice3 := make([]int, 2, 3)
	//mySlice3[7] = 10
	copy(mySlice, mySlice3)
	lol := append(mySlice3, 0)

	fmt.Println(mySlice3)
	fmt.Println(lol)
}
