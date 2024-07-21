package main

import "fmt"

func main() {
  tt := []struct {
    input    []int
    expected int
  }{
    {[]int{}, 0},
    {nil, 0},
    {[]int{10}, 10},
    {[]int{20, 30, 10, 50}, 10},
    {[]int{30, 50, 10}, 10},
  }

  for _, test := range tt {
    value, err := Min(test.input)
    if err != nil {
      fmt.Println(err)
      continue
    }
    fmt.Printf("Input: %d, Value: %d, Expected: %d, Match: %v\n", test.input, value, test.expected, value == test.expected)
  }
}

func Min(n []int) (int, error) {
  if len(n) == 0 {
    return 0, fmt.Errorf("slice %#v has no elements", n)
  }
  if len(n) == 1 {
    return n[0], nil
  }
  min := n[0]
  for _, num := range n[1:] {
    if num < min {
      min = num
    }
  }
  return min, nil
}
