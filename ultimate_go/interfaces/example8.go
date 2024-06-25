package main

import "fmt"

func main() {
  fmt.Println("Hello, World")
  fmt.Println(12345)
  fmt.Println(3.141519)
  fmt.Println(true)

  
  myPrintln("Hello, World")
  myPrintln(12345)
  myPrintln(3.141519)
  myPrintln(true)
}

func myPrintln(a interface{}) {
  switch v := a.(type) {
  case string:
    fmt.Printf("Is string  : type(%T) : value(%s)\n", v, v)
  case int:
    fmt.Printf("Is int     : type(%T) : value(%d)\n", v, v)
  case float64:
    fmt.Printf("Is float64 : type(%T) : value(%f)\n", v, v)
  default:
    fmt.Printf("Is unknown : type(%T) : value(%v)\n", v, v)
  }
}
