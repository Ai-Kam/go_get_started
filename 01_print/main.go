package main

import (
	"fmt"
)

func main() {
	fmt.Println("Printf:")
	x := "Hello"
	fmt.Printf("%s\n", x)
	pi := 3.14159265
	fmt.Printf("Pi = %f\n", pi)
        fmt.Println("Println:")
	fmt.Println(x, pi)
}
