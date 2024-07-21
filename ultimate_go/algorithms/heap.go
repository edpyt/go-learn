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

  heapSort(numbers)
  fmt.Println("Sequential:", numbers)
}

func heapSort(numbers []int) []int {
  for index := (len(numbers) / 2) - 1; index >= 0; index-- {
    numbers = moveLargest(numbers, len(numbers), index)
  }

  size := len(numbers)
  for index := size - 1; index >= 1; index-- {
    numbers[0], numbers[index] = numbers[index], numbers[0]
    size--
    numbers = moveLargest(numbers, size, 0)
  }

  return numbers
}

func moveLargest(numbers []int, size int, index int) []int {
  cmpIdx1, cmpIdx2 := 2*index + 1, 2*index + 2
  largestValueIdx := index

  if cmpIdx1 < size && numbers[cmpIdx1] > numbers[largestValueIdx] {
    largestValueIdx = cmpIdx1
  }

  if cmpIdx2 < size && numbers[cmpIdx2] > numbers[largestValueIdx] {
    largestValueIdx = cmpIdx2
  }

  if largestValueIdx != index {
    numbers[index], numbers[largestValueIdx] = numbers[largestValueIdx], numbers[index]
    numbers = moveLargest(numbers, size, largestValueIdx)
  }

  return numbers
}
