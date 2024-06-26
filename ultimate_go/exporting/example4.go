package main

import (
  "fmt"
)

//  .go package -----------
type User struct {
  // Экспортируемы
  Name     string
  ID       int

  // Не экспортируемы
  password string
}
// Обычно все поля должны быть либо экспортируемы
// либо не экспортируемы

// ------------------------

func main() {
  u := users.User{
    Name: "Chole",
    ID:   10,

    // FIXME: unknown field password in struct literal of type user.User
    password: "xxxx"
  }

  fmt.Printf("User: %#v\n", u)

}
