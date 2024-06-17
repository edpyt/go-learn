package main

import "fmt"

type user struct {
  name string
}

func main() {
  u, _ := createUser()
  fmt.Printf("%+v\n", *u)

  if _, err := createUser(); err != nil {
    fmt.Println(err)
    return
  }
}

func createUser() (*user, error) {
  return &user{"Bob"}, nil
}
