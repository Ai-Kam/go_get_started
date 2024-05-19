package main

import (
	"fmt"
)

func main() {
	var m1 int
	m1 = 3
	fmt.Println(m1) // 3

	var m2 int = 3
	fmt.Println(m2) // 3

	var x1 = 3
	fmt.Printf("%v [%T]\n", x1, x1) // 3 [int]

	x2 := 3 // var x = 3 と同じ意味
	fmt.Printf("%v [%T]\n", x2, x2) // 3 [int]

	var x3, y1 int = 1, 3
	fmt.Printf("%v, %v\n", x3, y1) // 1, 3

	var x4, y2 = 1, 3
	fmt.Printf("%v, %v\n", x4, y2) // 1, 3

	x5, y3 := 1, 3 // var x, y int = 1, 3 と同じ
	fmt.Printf("%v, %v\n", x5, y3) // 1, 3

	x6, y4, z1 := 1, 3.14, "Hello!"
	fmt.Printf("%v, %v, %v\n", x6, y4, z1) // 1, 3.14, Hello!

	var (
		x7 = 1
		y5 = 3.14
		z2 = "Hello!"
	)
	fmt.Printf("%v, %v, %v\n", x7, y5, z2) // 1, 3.14, Hello!
}
