package main

import "fmt"

func main() {
	const M1 = 10

	var i1 int = M1 // int に代入
	var x1 float64 = M1 // float64 に代入
	fmt.Printf("%v [%T]\n", i1, i1) // 10 [int]
	fmt.Printf("%v [%T]\n", x1, x1) // 10 [float64]

	const (
		M2 = 10
		N1 = 123
	)
	var i2 int = M2 // int に代入
        var x2 float64 = N1 // float64 に代入
        fmt.Printf("%v [%T]\n", i2, i2) // 10 [int]
        fmt.Printf("%v [%T]\n", x2, x2) // 123 [float64]

	const M3 int = 10
	var i3 int = M3 // これは OK
//	var x float64 = M // エラー
	// cannot use M (type int) as type float64 in assignment
	var x3 float64 = float64(M3)
        fmt.Printf("%v [%T]\n", i3, i3) // 10 [int]
        fmt.Printf("%v [%T]\n", x3, x3) // 10 [float64]

	const (
		north1 = iota
		south1
		east1
		west1
	)
	fmt.Printf("%v %v %v %v\n", north1, south1, east1, west1)
	// 0 1 2 3

        const (
                north2 = iota + 10
                south2
                east2
                west2
        )
        fmt.Printf("%v %v %v %v\n", north2, south2, east2, west2)
        // 10 11 12 13
}
