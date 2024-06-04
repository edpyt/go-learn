package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0

	for ; z <= x ; {
		// z -= (z * z - x) / (2 * z)
		z *= 2
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
