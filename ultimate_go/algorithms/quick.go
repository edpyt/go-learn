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

  QuickSort(numbers)
  fmt.Println("Sequential:", numbers)
}

func QuickSort(numbers []int) []int {
  return quickSort(numbers, 0, len(numbers) - 1)
}

func quickSort(numbers []int, leftIdx, rightIdx int) []int {
  switch {
  case leftIdx > rightIdx:
    return numbers

  case leftIdx < rightIdx:
    numbers, pivotIdx := partition(numbers, leftIdx, rightIdx)

    quickSort(numbers, leftIdx, pivotIdx-1)
    quickSort(numbers, pivotIdx + 1, rightIdx)
  }
  return numbers
}

func partition(numbers []int, leftIdx, rightIdx int) ([]int, int) {
  pivot := numbers[rightIdx]

  for smallest := leftIdx; smallest < rightIdx; smallest++ {
    if numbers[smallest] < pivot {
      numbers[smallest], numbers[leftIdx] = numbers[leftIdx], numbers[smallest]
      leftIdx++
    }
  }

  numbers[leftIdx], numbers[rightIdx] = numbers[rightIdx], numbers[leftIdx]

  return numbers, leftIdx
}
