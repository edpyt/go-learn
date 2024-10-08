package main

import (
  "fmt"
  "runtime"
  "sync"
)

func init() {
  runtime.GOMAXPROCS(2)
}

func main() {
  var wg sync.WaitGroup
  wg.Add(2)

  fmt.Println("Start Goroutines")

  go func() {
    for count := 0; count < 3; count++ {
      for r := 'a'; r <= 'z'; r++ {
        fmt.Printf("%c ", r)
      }
    }

    wg.Done()
  }()

  go func() {
    for count := 0; count < 3; count++ {
      for r := 'A'; r <= 'Z'; r++ {
        fmt.Printf("%c ", r)
      }
    }

    wg.Done()
  }()

  fmt.Println("Waiting To Finish")
  wg.Wait()

  fmt.Println("\nTerminating Program")
}
