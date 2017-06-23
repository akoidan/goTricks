package main

import "fmt"

func main() {
	appendToAll := 2
	funInside := func(a int, b int) int {
		return a + b + appendToAll
	}
	fmt.Println(funInside( 2,4))
}