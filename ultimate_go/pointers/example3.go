package main

import "fmt"

type user struct {
  name    string
  email   string
  logins  int
}

func main() {
  bill := user{
    name:      "Bill",
    email:     "bill@ardanlabs.com",
  }

  // Передаём адрес значения bill
  display(&bill)
  
  // Передаём адрес значения поля logins из значения bill
  increment(&bill.logins)

  display(&bill)
}

func increment(logins *int) {
  *logins++
  fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n\n", &logins, logins, *logins)
}

func display(u *user) {
  fmt.Printf("%p\t%+v\n", u, *u)
  fmt.Printf("Name: %q Email: %q Logins %d\n\n", u.name, u.email, u.logins)
}
