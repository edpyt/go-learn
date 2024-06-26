package main

import "fmt"


type user struct {
  name    string
  email   string
}

func (u *user) notify() {
  fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

type admin struct {
  person user
  level  string
}

func main() {
  ad := admin{
    person: user{
      name: "john smith",
      email: "john@yahoo.com",
    },
    level: "super",
  } 

  ad.person.notify()
}
