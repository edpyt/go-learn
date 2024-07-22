package main

import (
  "fmt"
  "math/rand"
)

func main() {
  numbers := []int{4, 42, 80, 83, 121, 137, 169, 182, 185, 180}
  find := rand.Intn(10)

  fmt.Println("Numbers", numbers)
  fmt.Println("Find.  :", numbers[find])

  idx, err := binarySearchIterative(numbers, numbers[find])
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("Found   : Idx", idx)
}

func binarySearchIterative(sortedList []int, target int) (int, error) {
  var leftIdx int
  rightIdx := len(sortedList) - 1

  for leftIdx <= rightIdx {
    mid := (leftIdx + rightIdx) / 2

    value := sortedList[mid]

    switch {
    case value == target:
      return mid, nil
    case value > target:
      rightIdx = mid -1
    case value < target:
      leftIdx = mid + 1
    }
  }

  return -1, fmt.Errorf("target not found")
}
