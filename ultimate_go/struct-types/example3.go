package main

import "fmt"

type example struct {
	flag      bool
	counter   int16
	pi        float32
}

func main() {
	e := struct {
		flag     bool
		counter  int16
		pi       float32
	}{
		flag:    true
		counter: 10
		pi:      3.141592
	}

	var ex example

	ex = e

	fmt.Printf("%+v\n", ex)
	fmt.Printf("%+v\n", e)
	fmt.Printf("Flag", e.flag)
	fmt.Printf("Counter", e.counter)
	fmt.Println("Pi", e.pi)
}
