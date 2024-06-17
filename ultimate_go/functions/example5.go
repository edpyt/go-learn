// Восстановление после паники
package main

import (
  "fmt"
  "runtime"
)

func main() {
  if err := testPanic(); err != nil {
    fmt.Println("Error:", err)
  }
}

func testPanic() (err error) {
  defer catchPanic(&err)

  err = mimicError("1")

  var p *int
  *p = 10

  fmt.Println("End Test")
  return err
}

func catchPanic(err *error) {
  if r := recover(); r != nil {
    fmt.Println("PANIC Deferred")

    buf := make([]byte, 10000)
    runtime.Stack(buf, false)
    fmt.Println("Stack Trace:", string(buf))

    if err != nil {
      *err = fmt.Errorf("%v", r)
    }
  }
}

func mimicError(key string) error {
  return fmt.Errorf("Mimic Error : %s", key)
}
