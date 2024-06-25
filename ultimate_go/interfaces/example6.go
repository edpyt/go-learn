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

func main() {
  svc := userSVC{
    host: "localhost:3434",
  }

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

  svc   := f.(*userSVC)
  log.Println("queried", svc.host)

  return nil
}
