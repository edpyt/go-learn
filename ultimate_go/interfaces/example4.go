package main

import "fmt"

type duration int

func (d *duration) notify() {
  fmt.Println("Sending Notification in", *d)
}

func main() {
  duration(42).notify()
}
