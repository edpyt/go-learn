package main

import (
  "fmt"
  "reflect"
)

type UnmarshalTypeError struct {
  Value string
  Type  reflect.Type
}

func (e *UnmarshalTypeError) Error() string {
  return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String() 
}
