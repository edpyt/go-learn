package main

import "fmt"

func main() {
  tt := []string{"", "G", "bob", "otto", "konichiwa", "test"}

  for _, input := range tt {
    success := IsPalindrome[string](input)

    switch success {
    case true:
      fmt.Printf("%q is a palindrome\n", input)
    case false:
      fmt.Printf("%q is NOT a palindrome\n", input)
    }
  }

  tt1 := []int{-1, 1, 9, 10, 1001, 125}

  for _, input := range tt1 {
    success := IsPalindrome[int](input)

    switch success {
    case true:
      fmt.Printf("%d is a palindrome\n", input)
    case false:
      fmt.Printf("%d is NOT a palindrome\n", input)
    }
  }
}

func IsPalindrome[T ~string | ~int](input T) bool {
  switch p := any(input).(type) {
  case string:
    if p == "" || len(p) == 1 {
      return true
    }

    runes := []rune(p)

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
      if runes[i] != runes[j] {
        return false
      }
    }

    return true

  case int:
    if p < 0 {
      return false
    }

    if p >= 0 && p < 10 {
      return true
    }

    rev := func(num int) int {
      var result int

      for num != 0 {
        last := num % 10
        result = result * 10
        result += last
        num = num / 10
      }

      return result
    }(p)

    return p == rev

  default:
    return false
  }
}
