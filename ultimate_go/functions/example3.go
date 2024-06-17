// Повторное объявление
package main

import "fmt"

type user struct {
  id int
  name string
}

func main() {
  var err1 error

  u, err1 := getUser()
  if err1 != nil {
    return
  }

  fmt.Println(u)

  u, err2 := getUser()
  if err2 != nil {
    return
  }

  fmt.Println(u)
}

func getUser() (*user, error) {
  return &user{1432, "Betty"}, nil
}
