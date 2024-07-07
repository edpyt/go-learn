package main

import (
  "context"
  "fmt"
  "sync"
)

type myKey int

const key myKey = 0

func main() {
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  var wg sync.WaitGroup
  wg.Add(10)

  for i:= 0; i < 10; i++ {
    go func(id int) {
      defer wg.Done()

      ctx := context.WithValue(ctx, key, id)

      <-ctx.Done()
      fmt.Println("Cancelled:",id)
    }(i)
  }

  cancel()
  wg.Wait()
}
