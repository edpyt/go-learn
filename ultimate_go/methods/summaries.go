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

// Methods - is just functions
type data struct {
  name    string
  age     int
}

func (d data) displayName() {
  fmt.Println("My Name Is", d.name)
}

func (d *data) setAge(age int) {
  d.age = age
  fmt.Println(d.name, "Is Age", d.age)
}

func main() {
  // u1 := user{
  //   name: "bob",
  //   email: "bob.avenue@mail.com",
  // }
  //
  // u1.notify()
  // u1.changeEmail("new.email.bob.avenue@mail.com")
  //
  // d := data{
  //   name: "Bill",
  // }
  //
  // d.displayName()
  // d.setAge(21)
  //
  // // Вызов методов как функций
  // data.displayName(d)
  // (*data).setAge(&d, 21)

  d := data {
    name: "Bill",
  }

  f1 := d.displayName
  f1()
  d.name = "Joan"
  f1()

  fmt.Println()

  f2 := d.setAge
  f2(45)
  d.name = "Sammy"
  f2(45)
}
