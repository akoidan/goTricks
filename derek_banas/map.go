package main

import "fmt"

func main() {
	presAge := make(map[string] int)
	presAge["TheodoreRoosevelt"] = 3
	presAge["ff"] = 5
	fmt.Println(presAge)
	fmt.Println(len(presAge))
	delete(presAge, "ff")
	fmt.Println(len(presAge))
	i := []int{2,3,4}
	fmt.Printf("Total summ is %d\n", addThem(i))
	a, b := next2val(2)
	fmt.Printf("A %d, B %d\n",a,b)
	fmt.Printf("addThemAll %d",addThemAll(2,3,4,54,5,6))
}

func addThem(number []int) int {
	sum := 0
	for _, v := range number {
		sum += v
	}
	return sum
}


func next2val(number int) (int, int) {
	return number +1, number+2
}

func addThemAll(args ...int) (a int) {
	for _, v := range args {
		a += v
	}
	return
}