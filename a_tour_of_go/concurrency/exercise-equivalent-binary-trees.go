package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)

    for {
        n1, ok1 := <- ch1
        n2, ok2 := <- ch2
        if ok1 != ok2 || n1 != n2 {
        	return false
        }
        if !ok1 {
        	break;
        }
    }
	return true
}

func TestWalk() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
