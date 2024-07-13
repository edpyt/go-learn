package main

import (
  "sync"
  "math/rand"
)

const goroutines = 100


func main() {
  ch := make(chan int, goroutines)

  var wg sync.WaitGroup

  wg.Add(goroutines)

  for i := 0; i < goroutines; i++ {
    go func() {
      defer wg.Done()
      rndm := rand.Intn(1000)
      if rndm % 2 == 0 {
        return
      }
      ch <- i
    }()
  }

  go func() {
    wg.Wait()
    close(ch)
  }()

  availableGoroutines := goroutines

  var nums []int
  for availableGoroutines > 0 {
    nums = append(nums, <-ch)
    availableGoroutines--
  }

  println(nums)
}
