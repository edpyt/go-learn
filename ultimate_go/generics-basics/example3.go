package main

import (
  "reflect"
)

func printReflect(v interface{}) {
  print("Reflect: ")

  val := reflect.ValueOf(v)
  if val.Kind() != reflect.Slice {
    return
  }

  for i := 0; i < val.Len(); i++ {
    print(val.Index(i).Interface(), " ")
  }

  print("\n")
}

func main() {
  numbers := []int{1,2,3}
  printReflect(numbers)
  print(numbers)

  strings := []string{"A", "B", "C"}
  printReflect(strings)
}
