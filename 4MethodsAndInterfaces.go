package main
//TODO https://tour.golang.org/methods/2
import "fmt"

type XYVertex struct {
	x int
	y int
}

func (v XYVertex) Abs() int {
	return v.x*v.x + v.y*v.y
}

func main() {
	var v XYVertex = XYVertex{2,3}
	fmt.Println(v.Abs())
}
