package main

import (
  "fmt"
  "sort"
)

func main() {
  tt := []struct {
    input1 string
    input2 string
  }{
    {"", ""},
    {"god", "dog"},
    {"god", "do"},
    {"1001", "0110"},
  }

  for _, test := range tt {
    success := IsPermutation(test.input1, test.input2)

    switch success {
    case true:
      fmt.Printf("%q and %q is a permutation\n", test.input1, test.input2)
    case false:
      fmt.Printf("%q and %q is NOT a permutation\n", test.input1, test.input2)
    }
  }
}

type RuneSlice []rune

func (p RuneSlice) Len() int { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int) () { p[i], p[j] = p[j], p[i] }

func IsPermutation(str1, str2 string) bool {
  if len(str1) != len(str2) {
    return false
  }
  s1 := []rune(str1)
  s2 := []rune(str2)

  sort.Sort(RuneSlice(s1))
  sort.Sort(RuneSlice(s2))

  return string(s1) == string(s2)
}
