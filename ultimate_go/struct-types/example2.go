package main

import "fmt"

func main() {
	var e1 struct {
		flag 			bool
		counter   int16	
		pi        float32
	}

	fmt.Printf("%+v\n", e1)

	e2 := struct {
		flag      bool
		counter   int16
		pi        float32
	}{
		flag:     true,
		counter:  10,
		pi:       3.141592,
	}

	fmt.Printf("%+v\n",e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
