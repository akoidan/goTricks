package main

//TODO https://tour.golang.org/methods/6
import "fmt"

type XYVertex struct {
	x int
	y int
}

type lol float64


type MyFloat float64

func (v MyFloat) plusPlus(to MyFloat) MyFloat {
	v += to // haha doesn't change initial value
	return v
}

func (v *MyFloat) incInitial(to MyFloat) {
	*v += to
}


func incInitial(v *MyFloat, to MyFloat) {
	*v += to
}

func (v XYVertex) Abs() int {
	return v.x*v.x + v.y*v.y
}

func Abs(v XYVertex) int {
	return v.x*v.x + v.y*v.y
} // same as above


func testMethds() {
	var v XYVertex = XYVertex{2, 3}
	f := MyFloat(10)
	incInitial(&f, MyFloat(44))
	d := f.plusPlus(MyFloat(5))
	f.incInitial(MyFloat(9))
	fmt.Println(v.Abs(), Abs(v), d, f)
}

func main() {
	testMethds()
}