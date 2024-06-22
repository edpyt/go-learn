package main

import "fmt"

type user struct {
  name    string
  surname string
}

func main() {
  users := map[string]user{
    "Roy":    {"Rob", "Roy"},
    "Ford":   {"Henry", "Ford"},
    "Mouse":  {"Mickey", "Mouse"},
    "Jackson": {"Michael", "Jackson"},
  }

  for key, value := range users {
    fmt.Println(key, value)
  }

  println()

  for key := range users {
    fmt.Println(key)
  }
}
