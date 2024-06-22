package main

import "fmt"

type user struct {
  name string
  surname string
}

type users []user

func main() {
  // invalid map key type
  // ключом могут быть только хэшируемые типы
  // (которые могут сравниваться между собой)
  u := make(map[users]int)

  for key, value := range u {
    fmt.Println(key, value)
  }
}
