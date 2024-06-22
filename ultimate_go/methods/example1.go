package main

import "fmt"

type user struct {
  name string
  email string
}

func (u user) notify() {
  fmt.Printf("Send user email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
  u.email = email
  fmt.Printf("Changed User email to %s\n", email)
}

func main() {
  u1 := user{
    name: "bob",
    email: "bob.avenue@mail.com",
  }

  u1.notify()
  u1.changeEmail("new.email.bob.avenue@mail.com")
}
