package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error){
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var z float64 = 1.0

	for ; z <= x ; {
		z *= 2
	}
	
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
