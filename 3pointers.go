package main
// TODO https://tour.golang.org/moretypes/16

import (
	"fmt"
	"strings"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func basicPointer() {
	i := 42
	p := &i
	*p = 32 // TODO new
	fmt.Println(i)
}

type Vertex struct {
	x int
	y int
}

func testMap()  {
	//var m map[string]Vertex
	m := map[string]Vertex{
		"A": Vertex{2, 3},
		"B": {4, 6},
	}
	//m = make(map[string]Vertex, 10)
	m["dancehub"] = Vertex{3 , 4}
	delete(m, "A")
	a, ok := m["D"]
	fmt.Printf("Map %v?Value of a: %v, present %v", m, a, ok)

}

func wordCount(dd string) map[string]int {
	d := strings.Fields(dd)
	m := make(map[string]int, len(d))
	for _, v := range d {
		m[v] ++
	}
	return m
}

func acceptFuncAsArg(fn func(int, int) int, a int, b int) int {
	return fn(a, b)
}

func funValueTest() {
	fmt.Printf("2+3= %v\n", acceptFuncAsArg(sum, 2, 3))
	diff := func (a int, b int) int {
		return a-b
	}
	fmt.Printf("3-4= %v\n", acceptFuncAsArg(diff, 3, 4))
}

func closureFunc() func(x int) int{
	sum := 0
	inc := func(x int) int {
		sum += x
		return sum
	}
	return inc
}

func testClosureFunc () {
	pos, neg := closureFunc(), closureFunc()
	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-i*2))
	}
}

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a+b
		a = b
		b = c
		return c
	}
}

func fibonacciTest() {
	getNextFibonacci := fibonacci()
	for i:=0; i < 10; i++ {
		fmt.Println(getNextFibonacci())
	}
}

func wcTest() {
	wc.Test(wordCount)
}

func structTest() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v1 := Vertex{}
	fmt.Println(v1)
	p2 := &Vertex{y: 4}
	p2.x = 3e4 // same as (*p).x = 3
	fmt.Println(p2)
}

func arrayTest() (primes [5]int, a [3]string) {
	a[0] = "Hello"
	a[1] = "Word"
	a[0] = "Destroy"
	fmt.Printf("Array : %v , %v \n", a[0], a[1])
	fmt.Println(a)
	primes = [5]int{1, 2, 3, 5, 7}
	fmt.Printf("Primes: %v", primes)
	return
}

func slicesTest() {
	var a, b = arrayTest();
	_ = b
	var sliceA []int = a[1:4]
	sliceA = []int{2, 3, 4, 4, 5, 6, 7, 7, 7, 7, 7, }
	fmt.Println(sliceA)
}

func anotherSliceTest() {
	a := [8]int{1, 2, 3, 4, 5, 6}
	_ = [4]string{
		"B",
		"A",
		"C",
		"D",
	}
	b := a[2:4]
	b[1] = 100
	c := a[3:5]
	fmt.Printf(" type of b %T, and a %T", b, a)
	fmt.Printf("A: %v, C: %v", a, c)
}

func sliceStruct() {
	s := []struct {
		i int
		string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
	}
	//mySlice := s[2:]
	//mySlice2 := s[:]
	printSliceInfo(s)
	//fmt.Printf("My slice %v, slice2 %v", mySlice, mySlice2)
}

func printSliceInfo(mySlice []struct {
	i int
	string
}) {
	fmt.Println(mySlice[0].string)
	fmt.Printf("cap : %d, len %d, of slice :%v", cap(mySlice), len(mySlice), mySlice)
}

func printSimpleSliceInfo(mySlice []int) {
	fmt.Printf(" len %d, cap : %d, of slice :%v \n", len(mySlice), cap(mySlice), mySlice)
}

func sliceCapacityLenTest() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSimpleSliceInfo(s)
	// Slice the slice to give it zero length.
	s = s[:0]
	printSimpleSliceInfo(s)

	// Extend its length.
	s = s[:4]
	printSimpleSliceInfo(s)

	// Drop its first two values.
	s = s[2:]
	printSimpleSliceInfo(s)
}

func zeroSlice() {
	mySlice := []int{}
	var mySlice2 []int
	fmt.Println("MySlice : %v, cap: %v, len: %v", mySlice, cap(mySlice), len(mySlice))
	if mySlice == nil {
		fmt.Println("mySlice's nil")
	}
	if mySlice2 == nil {
		fmt.Println("mySlice2's nil")
	}
}

func arrayAutoCount() {
	b := [...]string{"a", "b", "c"}
	fmt.Printf("Array's length %v", len(b))
}

func builtinMake() {
	a := make([]int, 5)
	a[1] = 1
	a[2] = 2
	a[4] = 4
	printSimpleSliceInfo(a)
	b := a[1:3]
	printSimpleSliceInfo(b)
	c := b[3:4]
	printSimpleSliceInfo(c)
	//b = b[0:5]
	//b[5] = 4
	//fmt.Println(b)
}

func multiDimensionSlice() {
	ticTacToe := [][]string{
		{"_", "_", "O"},
		{"_", "X", "_"},
		{"_", "_", "_"},
	}
	resTic := append(ticTacToe, []string{"what", "new", "row?"})
	//ticTacToe = ticTacToe[0:4] doesn't work resTic points to new array
	for i := 0; i < len(ticTacToe); i++ {
		fmt.Println(strings.Join(ticTacToe[i], " "))
	}
	fmt.Printf("new slice %v", resTic)

}

func getPic(dx,dy int) (result [][]uint8){
	result = make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		result[x] = make([]uint8, dy)
		for y :=0; y < dy; y++ {
			if x > 0 && y > 0 && (y != x){
				result[x][y] = uint8((y + x) / (y - x) * y * y / x)
			} else {
				result[x][y] = uint8(x^y)
			}

		}
	}
	return
}

func drawPic() {
	pic.Show(getPic)
}

func sliceIterRange() {
	slice := []int{2,4,8,16,32}
	for i, v := range slice {
		fmt.Printf("%v*2=%v\n", i+1,v)
	}
}

func sum(a int, b int) int {
	return a+b
}



func main() {
	//basicPointer()
	//structTest()
	//slicesTest()
	//anotherSliceTest()
	//sliceStruct()
	//sliceCapacityLenTest()
	//zeroSlice()
	//builtinMake();
	//arrayAutoCount()
	//builtinMake()
	//multiDimensionSlice()
	//sliceIterRange()
	//drawPic()
	//testMap()
	//wcTest();
	//funValueTest()
	//testClosureFunc()
	fibonacciTest()
}

