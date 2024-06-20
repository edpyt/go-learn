package main

import "fmt"

type user struct {
  id int
  name string
}

func main() {
  u1 := user{
    id: 1432,
    name: "Betty",
  }

  u2 := user{
    id: 4367,
    name: "Janet",
  }

  display(u1, u2)

  u3 := []user{
    {24, "Bill"},
    {32, "Joan"},
  }

  display(u3...)

  change(u3...)
  fmt.Println("****************")
  for _, u := range u3 {
    fmt.Printf("%+v\n", u)
  }
}

func display(users ...user) {
  fmt.Println("**************")
  for _, u := range users {
    fmt.Printf("%+v\n", u)
  }
}

func change(users ...user) {
  users[1] = user{99, "Same Backing Array"}
}
