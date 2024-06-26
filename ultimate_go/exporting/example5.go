package main

import (
  "fmt"
)

//  .go package -----------
type user struct {
  // Экспортируемы
  Name     string
  ID       int
}

type Manager struct {
  Title      string
  user
}
// ------------------------

func main() {
  u := users.User{
    Name: "Chole",
    ID:   10,
  }

  u.Name = "Chole"
  u.ID = 10

  fmt.Printf("User: %#v\n", u)

}
