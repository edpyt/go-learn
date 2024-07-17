package main

import "fmt"

func stringifyAssert(v interface{}) []string {
  switch list := v.(type) {
  case []user:
    ret := make([]string, 0, len(list))
    for _, value := range list {
      ret = append(ret, value.String())
    }
    return ret
  
  case []customer:
    ret := make([]string, 0, len(list))
    for _, value := range list {
      ret = append(ret, value.String())
    }
    return ret
  }

  return nil
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
  users := []user{
    {name: "Bill", email: "bill@ardanlabs.com"},
    {name: "Ale", email: "ale@whatever.com"},
  }

  s1 := stringifyAssert(users)

  println("users:", s1)

  customers := []customer {
    {name: "Google", email: "edpyt@google.com"},
    {name: "MSFT", email: "edpyt@msft.com"},
  }

  s2 := stringifyAssert(customers)

  println("customers:", s2)
}
