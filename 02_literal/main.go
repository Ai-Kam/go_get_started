package main

import (
	"fmt"
	"math"
)

func main() {
	x1 := 123
	fmt.Printf("%v [%T]\n", x1, x1)
	// 123 [int]

	x2 := 1_2_3
	fmt.Printf("%v [%T]\n", x2, x2)
	// 123 [int]

	x3 := 0b1100
	fmt.Printf("%v [%T]\n", x3, x3) // 12 [int]

	x4 := 0123
	fmt.Printf("%v [%T]\n", x4, x4) // 83 [int]

	x5 := 0xBEEF
	fmt.Printf("%v [%T]\n", x5, x5) // 48879 [int]

	x6 := 3.14
	fmt.Printf("%v [%T]\n", x6, x6) // 3.14 [float64]

	x7 := 1.23e-3
	fmt.Printf("%v [%T]\n", x7, x7) // 0.00123 [float64]

	x8 := 0x1.fp3
	fmt.Printf("%v [%T]\n", x8, x8) // 15.5 [float64]
	y1 := (1 + 15.0/16.0) * math.Pow(2, 3)
	fmt.Printf("%v\n", y1) // 15.5

	x9 := 123i
	fmt.Printf("%v [%T]\n", x9, x9) // (0+123i) [complex128]

	x10 := '\101'
	y2 := '\x41'
	fmt.Printf("%c %c\n", x10, y2) // A A

	x11 := '\u672C'
	y3 := '\U0001F604'
	fmt.Printf("%c [%T]\n", x11, x11) // 本 [int32]
	fmt.Printf("%c [%T]\n", y3, y3) // 😄 [int32]

	x12 := "Hello, world!"
	y4 := "\101\102\x43"
	fmt.Println(x12) // Hello, world!
	fmt.Println(y4) // ABC

	x13 := `Hello,
	World!`
	fmt.Println(x13)
	// Hello,
	// World!

	x14 := `Hello, world!
	\101\102\x43`
	fmt.Println(x14)
	// Hello, world!
	// \101\102\x43
}
