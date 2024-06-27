package main

import "fmt"

type Animal struct {
  Name      string
  IsMammal  bool
}

type Dog struct {
  Animal
  PackFactor int
}

func (d *Dog) Speak() {
  fmt.Printf(
    "Woof! My name is %s, it is %t I am a mammal with a pack factor of %d. \n",
    d.Name,
    d.IsMammal,
    d.PackFactor,
  )
}

type Cat struct {
  Animal
  ClimbFactor int
}

func (c *Cat) Speak() {
  fmt.Printf(
    "Meow! My name is %s, it is %t I am a mammal with a climb factor of %d. \n",
    c.Name,
    c.IsMammal,
    c.ClimbFactor,
  )
}

type Speaker interface {
  Speak()
}

func main() {
  // Если мы используем интерфейс, то мы можем определить общий набор
  // методов поведения, который возможно использоваться для группировки
  // разных типов данных
  speakers := []Speaker{
    &Dog{
      Animal: Animal{
        Name:     "Fido",
        IsMammal: true,
      },
      PackFactor: 5,
    },
    &Cat{
      Animal: Animal{
        Name:     "Milo",
        IsMammal: true,
      },
      ClimbFactor: 4,
    },
  }

  for _, speaker := range speakers {
    speaker.Speak()
  }
}

