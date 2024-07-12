package main

import (
  "sync"
)

func main() {
  ch := make(chan int, 10)

  var wg sync.WaitGroup
  wg.Add(2)

  go func() {
    defer wg.Done()
    goroutine(ch)
  }()

  go func() {
    defer wg.Done()
    goroutine(ch)
  }()

  ch<-0

  wg.Wait()
}

func goroutine(ch chan int) {
  for {
    d, ok := <- ch
    if !ok {
      return
    }
    println(d)
    if d == 10 {
      close(ch)
      return
    }
    ch<-(d+1)
  }
}
