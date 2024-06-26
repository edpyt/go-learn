package main

import (
  "fmt"
  "play.ground/counters"
)

func main() {
  counter := counters.AlertCounter(10)

  fmt.Printf("Counter: %d\n", counter)
}
