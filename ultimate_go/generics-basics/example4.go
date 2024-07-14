package main

import "fmt"

func print[T any](slice []T) {
  fmt.Print("Generic: ")

  for _, v := range slice {
    fmt.Print(v, " ")
  }

  fmt.Print("\n")
}

func main() {
  numbers := []int{1,2,3}
  print(numbers)

  strings := []string{"A", "B", "C"}
  print(strings)
}
