package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func forever(end int) int {
	i := 1
	for {
		if i > end {
			break
		}
		i += i
	}
	return i
}

func while(end int) (i int) {
	i = 1
	for i < end {
		i += i
	}
	return
}

func fullFor(end int) (j int) {
	for j = 1; j < end; j += j {
	}
	return j
}

func ifStatement(end float64) {
	if d := math.Pow(end, 2); d < 10 {
		print("It works")
	} else {
		print("It doesnt work")
	}
}

func sqrt(number int, precision int) float64  {
	z := 1.0
	for i := 1; i < precision; i++ {
		z = z - (z*z-float64(number))/(2*z)
	}
	return z
}

func isSaturdaySoon(){
	nowDay := time.Now().Weekday()
	diff := time.Saturday - nowDay
	switch {
	case diff < 2:
		print("good")
	case diff < 4:
		print("bad")
	default:
		print("Holly fuck :(")
	}
	
}

func switchStatement()  {
	switch os := runtime.GOOS; os {
	case "darwin":
		print("OS X")
	case "linux":
		print("Linux")
	default:
		print(os)
	}
}

func getNumber(arg int) int {
	fmt.Printf("Getting number %v\n", arg)
	return arg
}

func deferStatement(number int)  {
	fmt.Println("Preparing deffer")
	defer fmt.Printf("Deffer %v is called\n", getNumber(number))
	fmt.Println("Deffer is used")
}

func stackDeffer() {
	for i := 0; i< 4; i++ {
		defer deferStatement(i)
	}
	fmt.Println("All deffers are stacked")

}
func main() {
	fmt.Println("Starting program")
	//fmt.Println(fullFor(100))
	//fmt.Println(while(100))
	//fmt.Println(forever(100))
	//ifStatement(2.0)
	//fmt.Printf("Newton's  %v, Math: %v", sqrt(4,4), math.Sqrt(4))
	//switchStatement()

	//isSaturdaySoon()
	stackDeffer()
	fmt.Println("Main is finished")
}
