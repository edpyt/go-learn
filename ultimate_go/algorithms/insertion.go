package main

import (
  "fmt"
  "math/rand"
)

func generateList(n int) []int {
  numbers := make([]int, n)

  for i := 0; i < n; i++ {
    numbers[i] = rand.Intn(n * 20)
  }

  return numbers
}

func main() {
  numbers := generateList(5)
  fmt.Println("Before:", numbers)

  insertionSort(numbers)
  fmt.Println("Sequential:", numbers)
}

func insertionSort(numbers []int) {
  var n = len(numbers)

  for i := 1; i < n; i++ {
    j := i

    for j > 0 {
      if numbers[j - 1] > numbers[j] {
        numbers[j - 1], numbers[j] = numbers[j], numbers[j - 1]
      }
      j--
    }
  }
}
