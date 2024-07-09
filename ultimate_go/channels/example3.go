// Ожидание задачи
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  waitForTask()
}

func waitForTask() {
  ch := make(chan string)

  go func() {
    d := <-ch
    fmt.Println("child: recv'd signal :", d)
  }()

  time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
  ch <- "data"
  fmt.Println("parent : sent signal")

  time.Sleep(time.Second)

  fmt.Println("-------------------------------------------------")
}
