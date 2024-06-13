package main

import "fmt"

type example struct {
	flag 			bool
	counter   int16
	pi        float32
}

func main() {
	var e1 example
	
	fmt.Printf("%+v\n", e1)

	e2 := example{
		flag: 		true,
		counter:  10,
		pi:       3.141592,
	}

	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
