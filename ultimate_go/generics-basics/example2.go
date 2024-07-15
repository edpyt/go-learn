package main

import "fmt"

func printAssert(v interface{}) {
  fmt.Print("Assert: ")

  switch list := v.(type) {
  case []int:
    for _, num := range list {
      fmt.Print(num, " ")
    }
  case []string:
    for _, str := range list {
      fmt.Print(str, " ")
    }
  }

  fmt.Print("\n")
}

func main() {
  numbers := []int{1,2,3}
  printAssert(numbers)

  strings := []string{"A", "B", "C"}
  printAssert(strings)
}
