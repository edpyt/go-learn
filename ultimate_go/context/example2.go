package main

import (
  "context"
  "fmt"
  "time"
)

func main() {
  ctx, cancel = context.WithCancel(context.Background())
  defer cancel()

  go func() {
    select {
      case <- time.After(100 * time.Millisecond): fmt.Println("moving on")
      case <- ctx.Done():
      fmt.Println("work complete")
    }
  }()

  time.Sleep(50 * time.Millisecond)

  cancel()

  time.Sleep(time.Second)
}
