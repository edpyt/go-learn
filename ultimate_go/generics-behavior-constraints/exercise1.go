package main

import (
  "fmt"
  "encoding/json"
)

func marshal[T json.Marshaler](v T) ([]byte, error) {
  data, err :=  json.Marshal(v)
  if err != nil {
    return nil, err
  }
  return data, nil
}

type user struct {
  Name  string `json:"name"`
  Email string `json:"email"`
}

func (u user) MarshalJSON() ([]byte, error) {
  v := fmt.Sprintf("{\"name\": %q, \"email\": %q}", u.Name, u.Email)
  return []byte(v), nil
}

func main() {
  u := user{
    Name: "Bob",
    Email: "bob@ardanlabs.com",
  }

  serializedU, _ := marshal(u)

  fmt.Println(string(serializedU))
}
