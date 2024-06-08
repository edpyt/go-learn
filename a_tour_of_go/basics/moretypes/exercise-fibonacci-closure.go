package main

import "fmt"

func fibonacci() func() int {
	fib := 0
	prev := 1
	return func() int {
		nth := fib + prev
		prev = fib
		fib = nth
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
