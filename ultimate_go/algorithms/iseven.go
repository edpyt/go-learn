package main

import "fmt"

func main() {
  fmt.Println(8, ":", IsEven(8))
  fmt.Println(15, ":", IsEven(15))
  fmt.Println(4, ":", IsEven(4))
}

func IsEven(num int) bool {
  return num&1 == 0
}
