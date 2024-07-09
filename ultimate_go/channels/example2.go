// Веерный выход

package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  fanOut()
}

func fanOut() {
  children := 2000 
  ch := make(chan string, children)

  for c := 0; c < children; c++ {
    go func(child int) {
      time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
      ch <- "data"
      fmt.Println("child : sent signal :", child)
    }(c)
  }

  for children > 0 {
    d := <-ch
    children--
    fmt.Println(d)
    fmt.Println("parent : recv'd signal :", children)
  }

  time.Sleep(time.Second)

  fmt.Println("----------------------------------")
}
