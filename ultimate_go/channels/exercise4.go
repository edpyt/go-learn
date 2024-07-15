package main

import (
  "fmt"
  "runtime"
  "math/rand"
  "sync"
)

func main() {
  values := make(chan int)

  shutdown := make(chan struct{})

  poolSize := runtime.GOMAXPROCS(0)

  var wg sync.WaitGroup
  wg.Add(poolSize)

  for i := 0; i < poolSize; i++ {
    go func(id int) {
      for {
        n := rand.Intn(1000)

        select {
        case values <- n:
          fmt.Printf("Worker %d sent %d\n", id, n)
        case <-shutdown:
          fmt.Printf("Worker %d shutting down\n", id)
          wg.Done()
          return
        }
      }
    }(i)
  }

  var nums []int
  for i := range values {
    if i%2 == 0 {
      println("Discarding", i)
      continue
    }

    println("Keeping", i)
    nums = append(nums, i)

    if len(nums) == 100 {
      break
    }
  }

  println("Receiver sending shutdown signal")
  close(shutdown)

  wg.Wait()

  fmt.Printf("Result count: %d\n", len(nums))
  println(nums)
}
