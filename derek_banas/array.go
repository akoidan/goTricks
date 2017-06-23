package main

import "fmt"

func main() {
	var arr1[5] float32;
	arr1[0] = 23.2
	arr1[1] = 423.2
	arr1[2] = 55.2
	arr1[3] = 223.2
	arr1[4] = 2453.2
	fmt.Print(arr1)
	arr2 := [5]int {1,2,3,4}
	fmt.Println(arr2)

	for i, v := range arr2 { // value can be skipped
		// undersocre can be used to skip var
		fmt.Printf("i: %d, v %d", i, v)
	}
}