package main

import "fmt"

type Mover interface {
  Move()
}

type Locker interface {
  Lock()
}

type MoveLocker interface {
  Mover
  Locker
}

type bike struct{}

func (bike) Move() {
  fmt.Println("Moving the bike")
}

func (bike) Lock() {
  fmt.Println("Locking the bike")
}

func (bike) Unlock() {
  fmt.Println("Unlocking the bike")
}

func main() {
  var ml MoveLocker
  var m  Mover

  ml = bike{}

  m = ml

  // FIXME: cannot use m (type Mover) as type MoveLocker
  // ml = m

  b := m.(bike)
  ml = b
}
