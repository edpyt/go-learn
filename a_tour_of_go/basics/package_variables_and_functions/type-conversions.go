package main

import (
	"fmt"
	"math"
)

func typeConversions() {
//	i := 42
//	f := float64(i)
//	u := uint(f)

	var i = 42
	var f = float64(i)
	var u = uint(f)

	fmt.Printf("typeConversionValue: %v, u var: %v", f, u)
}

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*x))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	typeConversions()
}
