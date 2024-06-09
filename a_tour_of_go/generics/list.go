package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val T
}

func (xs *List[T]) String() string {
	if xs.next == nil {
		return fmt.Sprintf("{ val: %v, next: <nil> }", xs.val)
	} else {
		return fmt.Sprintf("{ val: %v, next: %s }", xs.val, fmt.Sprint(xs.next))
	}
} 

func main() {
	xs := new(List[int])
	xs.val = 5
	xs.next = nil
	fmt.Println(xs)
}
