package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var (
		y float64 = 1.0
		z float64 = 1.0
	) 

	for ; y <= x ; {
		z -= (z * z - x) / (2 * z)
		y *= 2
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
