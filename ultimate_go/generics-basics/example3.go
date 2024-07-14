package main

import (
  "fmt"
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
}
