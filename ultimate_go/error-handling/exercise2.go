package main

import "errors"

type appError struct {
  err error
  message string
  code    int
}

func (a appError) temporary() bool {
  if a.code == 9 {
    return false
  }
  return true
}

func main () {
  checkFlag(true)
}

func checkFlag(b bool) *appError {
  if b {
    return &appError{err: errors.New("error"), message: "app error", code: 9}
  } else {
    return &appError{err: errors.New("bob"), message: "ahaha", code: 10}
  }
}


