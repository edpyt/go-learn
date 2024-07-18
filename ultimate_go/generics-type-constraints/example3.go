package main

import "fmt"

type person struct {
  name  string
  email string
}

func (p person) match(v person) bool {
  return p.name == v.name
}

type food struct {
  name     string
  category string
}

func (f food) match(v food) bool {
  return f.name == v.name
}

type matcher[T any] interface {
  person | food
  match(v T) bool
}

func match[T matcher[T]](list []T, find T) int {
  for i, v := range list {
    if v.match(find) {
      return i
    }
  }
  return -1
}

func main() {
  people := []person{
    {name: "bill", email: "bill@mail.com"},
    {name: "ed", email: "ed@mail.com"},
    {name: "emil", email: "emil@mail.com"},
  }
  findPerson := person{name: "tony"}

  i := match(people, findPerson)
  fmt.Printf("Match: Idx: %d for %s\n", i, findPerson.name)

  foods := []food{
    {name: "apple", category: "fruit"},
    {name: "carrot", category: "veg"},
    {name: "chicken", category: "meat"},
  }
  findFood := food{name: "apple"}

  i = match(foods, findFood)
  fmt.Printf("Match: Idx: %d for %s\n", i, findFood.name)
}
