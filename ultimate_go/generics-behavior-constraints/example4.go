package main

import "fmt"

func stringify[T fmt.Stringer](slice []T) []string {
  ret := make([]string, 0, len(slice))

  for _, value := range slice {
    ret = append(ret, value.String())
  }

  return ret
}

type user struct {
  name  string
  email string
}

func (u user) String() string {
  return fmt.Sprintf("{type: \"user\", name: %q, email: %q}", u.name, u.email)
}

type customer struct {
  name  string
  email string
}

func (u customer) String() string {
  return fmt.Sprintf("{type: \"customer\", name: %q, email: %q}", u.name, u.email)
}

func main() {
  users := []user{
    {name: "Bill", email: "bill@ardanlabs.com"},
    {name: "Ale", email: "ale@whatever.com"},
  }

  s1 := stringify(users)

  println("users:", s1)

  customers := []customer{
    {name: "Google", email: "edpyt@google.com"},
    {name: "MSFT", email: "edpyt@msft.com"},
  }

  s2 := stringify(customers)

  println("customers:", s2)
}
