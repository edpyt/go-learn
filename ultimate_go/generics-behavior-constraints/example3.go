package main

import (
  "fmt"
  "reflect"
)

func stringifyReflect(v interface{}) []string {
  val := reflect.ValueOf(v)
  if val.Kind() != reflect.Slice {
    return nil
  }

  ret := make([]string, 0, val.Len())

  for i := 0; i < val.Len(); i++ {
    m := val.Index(i).MethodByName("String")
    if !m.IsValid() {
      return nil
    }

    data := m.Call(nil)
    ret = append(ret, data[0].String())
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
  users :=  []user{
    {name: "Bill", email: "bill@ardanlabs.com"},
    {name: "Ale", email: "ale@whatever.com"},
  }

  s1 := stringifyReflect(users)

  println("users:", s1)

  customers := []customer{
    {name: "Google", email: "edpyt@google.com"},
    {name: "MSFT", email: "edpyt@msft.com"},
  }

  s2 := stringifyReflect(customers)

  println("customers:", s2)
}
