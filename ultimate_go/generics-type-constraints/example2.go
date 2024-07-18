package main

import "fmt"

func index[T comparable](list []T, find T) int {
  for i,v := range list {
    if v == find {
      return i
    }
  }
  return -1
}

type person struct {
  name  string
  email string
}

func main() {
  durations := []int{5000, 10, 40}
  findDur := 10

  i := index(durations, findDur)
  fmt.Printf("Index: %d for %d\n", i, findDur)

  people := []person{
    {name: "bill", email: "bill@mail.com"},
    {name: "ed", email: "ed@mail.com"},
    {name: "emil", email: "emil@mail.com"},
  }
  findPerson := person{name: "tony", email: "tony@mail.com"}

  i = index(people, findPerson)
  fmt.Printf("Index: %d for %s\n", i, findPerson.name)
}
