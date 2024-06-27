package main

import (
  "fmt"
  "math/rand"
)

type car struct{}

func (car) String() string {
  return "Vroom!"
}

type cloud struct{}

func (cloud) String() string {
  return "Big Data!"
}

func main()  {
  mvs := []fmt.Stringer{
    car{},
    cloud{},
  }

  for i := 0; i  < 10; i++ {
    rn := rand.Intn(2)

    if v, is := mvs[rn].(cloud); is {
      fmt.Println("Got Lucky:",v)
    } else {
      fmt.Println("Got Unlucky")
    }
  }
}
