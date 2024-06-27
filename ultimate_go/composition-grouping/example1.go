package main

type Animal struct {
  Name      string
  IsMammal  bool
}

func (a *Animal) Speak() {
  fmt.Printf(
    "UGH! My name is %s, it is %t I am a mammal\n",
    a.Name,
    a.IsMammal,
  )
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

func main() {
  animals := []Animal{
    Dog{
      Animal: Animal{
        Name:     "Fido",
        IsMammal: true,
      },
      PackFactor:5,
    },
    Cat{
      Animal: Animal{
        Name:     "Milo",
        IsMammal: true,
      },
      ClimbFactor:4,
    },
  }

  for _, animal := range animals {
    animal.Speak()
  }
}
