package main

import (
	"fmt"
)

func main() {
	var s string = "Hello, world!"
	var s2 string = "Hello \"world\"!"

	fmt.Println(s)  // Hello, world!
	fmt.Println(s2) // Hello "world"!
}
