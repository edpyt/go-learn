package main

import (
  "fmt"
  "log"
)

type user struct {
  id    int
  name  string
}

type finder interface {
  find(id int) (*user, error)
}

type userSVC struct {
  host  string
}

func (*userSVC) find(id int) (*user, error) {
  return &user{id: id, name: "Anna Walker"}, nil
}

type mockSVC struct{}

func (*mockSVC) find(id int) (*user, error) {
  return &user{id: id, name: "Jacob Walker"}, nil
}

func main() {
  var svc mockSVC

  if err := run(&svc); err != nil {
    log.Fatal(err)
  }
}

func run(f finder) error {
  u, err :=  f.find(1234)
  if err != nil {
    return err
  }
  fmt.Printf("Found user %+v\n", u)

  if svc, ok := f.(*userSVC); ok {
    log.Println("queried", svc.host)
  }

  return nil
}
