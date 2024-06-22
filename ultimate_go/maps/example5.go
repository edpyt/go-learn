// Итерирование по map в алфавитном пордяке

package main

import (
  "fmt"
  "sort"
)

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

  var keys []string
  for key := range users {
    keys = append(keys, key)
  }
  
  sort.Strings(keys)

  for _, key := range keys {
    fmt.Println(key, users[key])
  }

}
