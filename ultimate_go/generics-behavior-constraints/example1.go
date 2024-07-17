package main

import "fmt"

func stringifyUsers(users []user) []string {
  ret := make([]string, 0, len(users))
  for _, user := range users {
    ret = append(ret, user.String())
  }
  return ret
}

func stringifyCustomers(customers []customer) []string {
  ret := make([]string, 0, len(customers))
  for _, customer := range customers {
    ret = append(ret, customer.String())
  }
  return ret
}

type user struct {
  name string
  email string
}

func (u user) String() string {
  return fmt.Sprintf("{type: \"user\", name: %q, email: %q}", u.name, u.email)
}

type customer struct {
  name string
  email string
}

func (u customer) String() string {
  return fmt.Sprintf("{type: \"customer\", name: %q, email: %q}", u.name, u.email)
}

func main() {
  users := []user {
    {name: "Bill", email: "bill@adasje.com"},
    {name: "Ale", email: "ale@adasje.com"},
  }

  s1 := stringifyUsers(users)

  println("users:", s1)

  customers := []customer{
    {name: "Google", email: "edpyt@google.com"},
    {name: "MSFT", email: "edpyt@msft.com"},
  }

  s2 := stringifyCustomers(customers)

  println("customers:", s2)
}
