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
  numbers := generateList(10)
  fmt.Println("Before:", numbers)

  bubbleSort(numbers)
  fmt.Println("Sequential:", numbers)
}

func bubbleSort(numbers []int) {
  n := len(numbers)
  
  for i := 0; i < n; i++ {
    if !sweep(numbers, i) {
      return
    }
  }
}

func sweep(numbers []int, currentPass int) bool {
  var idx int
  var swap bool

  idxNext := idx + 1
  n := len(numbers)

  for idxNext < (n - currentPass) {
    a := numbers[idx]
    b := numbers[idxNext]

    if a > b {
      numbers[idx] = b
      numbers[idxNext] = a
      swap = true
    }

    idx++
    idxNext = idx + 1
  }

  return swap
}

