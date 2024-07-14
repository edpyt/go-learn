package main

import "fmt"

func printNumbers(numbers []int) {
  fmt.Print("Numbers: ")

  for _, num := range numbers {
    fmt.Print(num, " ")
  }

  fmt.Print("\n")
}

func printStrings(strings []string) {
  fmt.Print("Strings: ")

  for _, str := range strings {
    fmt.Print(str, " ")
  }

  fmt.Print("\n")
}

func main() {
  numbers := []int{1,2,3}
  printNumbers(numbers)

  strings := []string{"A", "B", "C"}
  printStrings(strings)
}
