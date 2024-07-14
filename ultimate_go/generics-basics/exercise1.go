package main

import "encoding/json"

func marshal[T any](u interface{}) []byte {
  b, err := json.Marshal(u)

  if err != nil {
    println(err)
    return nil
  }

  return b
}

type User struct {
  Name string
  Age int
}

func main() {
  u := User{Name: "Bob", Age: 21}

  res := marshal[interface{}](u)

  print(string(res))
}
