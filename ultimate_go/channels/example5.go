// Веерный выход с семафором
package main

import (
  "fmt"
  "math/rand"
  "runtime"
  "time"
)

func main() {
  fanOutSem()
}

func fanOutSem() {
  children := 2000
  ch := make(chan string, children)

  g := runtime.GOMAXPROCS(0)
  sem := make(chan bool, g)

  for c := 0; c < children; c++ {
    go func(child int) {
      sem <- true
      {
        t := time.Duration(rand.Intn(200)) * time.Millisecond
        time.Sleep(t)
        ch <- "data"
        fmt.Println("child : sent signal :", child)
      }
      <-sem
    }(c)
  }

  for children > 0 {
    d := <-ch
    children--
    fmt.Println(d)
    fmt.Println("parent : recv'd signal :", children)
  }

  time.Sleep(time.Second)

  fmt.Println("---------------------------------------------------------")
}
