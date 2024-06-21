package main

import "fmt"

type user struct {
  name    string
  surname string
}

func main() {
  users := make(map[string]user)

  users["Roy"] = user{"Rob", "Roy"}
  users["Ford"] = user{"Henry", "Ford"}
  users["Mouse"] = user{"Mickey", "Mouse"}
  users["Jackson"] = user{"Michael", "Jackson"}

  mouse := users["Mouse"]

  fmt.Printf("%+v\n", mouse)

  users["Mouse"] = user{"Jerry", "Mouse"}

  fmt.Printf("%+v\n", users["Mouse"])

  fmt.Println(len(users))
  delete(users, "Roy")
  fmt.Println(len(users))

  delete(users, "Roy")
  fmt.Println("Goodbye.")
}
