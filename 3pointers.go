package main
// TODO https://tour.golang.org/moretypes/6
import (
	"fmt"
)

func basicPointer() {
	i := 42
	p := &i
	*p = 32
	fmt.Println(i)
}

type Vertex struct {
	x int
	y int
}

func structTest() {
	v := Vertex{1,2}
	fmt.Println(v)
	v1 := Vertex{}
	fmt.Println(v1)
	p2 := &Vertex{y:4}
	p2.x = 3e4 // same as (*p).x = 3
	fmt.Println(p2)
}

func main() {
	basicPointer()
	structTest()
}