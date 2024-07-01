package main

import (
  "errors"
  "fmt"
)

var (
  ErrInvalidValue = "Invalid Value Error"
  ErrAmountTooLarge = "Amount too large!"
)

func checkAmount(amount float64) error {
  if amount == 0 {
    return errors.New(ErrInvalidValue)
  } else if amount > 1000 {
    return errors.New(ErrAmountTooLarge)
  } else {
    return nil
  }
}

func main() {
  am := 1001.0
  
  if err := checkAmount(am); err != nil {
    fmt.Printf("Error! %s", err)
  }
}
