package main

// difference := = var
import (
//"math/cmplx"
)
import "fmt"

var java, python bool

const (
	a = 1 << 100
	b = a >> 97
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func add(x, y int) (sum int, difference int) {

	//println(i)
	sum = x + y
	difference = x - y
	return
}

func main() {
	var a , b  = add(2, 3)
	//var i = 344333333333333333333333333333333333.3
	fmt.Println(needInt(b))
	fmt.Println(needFloat(b))
	fmt.Println(a)
}
