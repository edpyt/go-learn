package main

import "fmt"

func main() {
  friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

  for i,v := range friends {
    fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &friends[i])
  }
}
