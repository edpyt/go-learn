package main

import (
  "fmt"
  "math/rand"
)

func main() {
  numbers := []int{4, 42, 80, 83, 121, 137, 169, 182, 185, 180}
  find := rand.Intn(10)

  fmt.Println("Numbers:", numbers)
  fmt.Println("Find   :", numbers[find])

  idx, err := binarySearchRecursive(numbers, numbers[find], 0, len(numbers))
  if err != nil {
    fmt.Println(err)
    return 
  }

  fmt.Println("Found  : Idx", idx)
}

func binarySearchRecursive(sortedList []int, target int, leftIdx, rightIdx int) (int, error) {
  midIdx := (leftIdx + rightIdx) / 2

  if leftIdx <= rightIdx  {
    switch{
    case sortedList[midIdx] == target:
      return midIdx, nil
    case sortedList[midIdx] > target:
      return binarySearchRecursive(sortedList, target, leftIdx, midIdx - 1)
    case sortedList[midIdx] < target:
      return binarySearchRecursive(sortedList, target, midIdx + 1, rightIdx)
    }
  }

  return -1, fmt.Errorf("target not found")
}

