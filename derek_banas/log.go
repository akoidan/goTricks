package main

import "fmt"

func main() {
	var a = true;
	var b = false;
	print(a || b)

	i := 1;
	for i < 10 {
		i++;
		fmt.Println(i)
	}

	fmt.Printf("true  != false: %t \n", true != false)


	for j:=5; j < 8; j++ {
		fmt.Println(j)
		if j == 6 {
			continue
		} else if j == 7 {

		} else {
			fmt.Println("else")
		}
	}

	k := 99;
	switch k  {
		case 16: fmt.Println(k)
		case 14: fmt.Println(k+1)
	}
	fmt.Println("-------------------------")

}