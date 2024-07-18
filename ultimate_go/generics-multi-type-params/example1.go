package main

import "fmt"

func Print[L any, V fmt.Stringer](labels []L, vals []V) {
  for i, v := range vals {
    fmt.Println(labels[i], v.String())
  }
}

type user struct {
  name string
}

func (u user) String() string {
  return u.name
}

func main() {
  labels := []int{1,2,3}
  names := []user{{"bill"}, {"jill"}, {"joan"}}

  Print(labels, names)
}
