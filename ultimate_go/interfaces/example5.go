package main

import "fmt"

type printer interface {
  print()
}

type cannon struct {
  name string
}

func (c cannon) print() {
  fmt.Printf("Printer Name: %s\n", c.name)
}

type epson struct {
  name string
}

func (e *epson) print()  {
  fmt.Printf("Printer Name: %s\n", e.name)
}

func main() {
  c := cannon{"PIXMA TR4520"}
  e := epson{"WorkForce Pro WF-3720"}

  printers := []printer{
    c,
    &e,
  }

  c.name = "PROGRAF PRO-1000"
  e.name = "Home XP-4100"

  for _, p := range printers {
    p.print()
  }
}
