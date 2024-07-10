// NOTE: Ограниченный пул работ

package main

import (
  "fmt"
  "time"
)

func main() {
  drop()
}

func drop() {
  const cap = 100
  ch := make(chan string, cap)

  go func() {
    for p := range ch {
      fmt.Println("child : recv'd signal :", p)
    }
  }()

  const work = 2000
  for w := 0; w < work; w++ {
    select {
    case ch <- "data":
      fmt.Println("parent : sent signal :", w)
    default:
      fmt.Println("parent : dropped data :", w)
    }
  }

  close(ch)
  fmt.Println("parent : sent shutdown signal")

  time.Sleep(time.Second)
  fmt.Println("------------------------------------")
}
