package main

import "fmt"

type addOnly interface {
  string | int | int8 | int16 | int32 | int64 | float64
}

func Add[T addOnly](v1 T, v2 T) T {
  return v1 + v2
}

func main() {
  fmt.Println(Add(10, 20))
  fmt.Println(Add("A", "B"))
  fmt.Println(Add(3.14159, 2.96))
}
