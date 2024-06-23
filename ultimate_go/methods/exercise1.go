package main

type baseballPlayer struct {
  name   string
  atBats float64
  hits   float64
}

func (b baseballPlayer) average() float64 {
  return b.hits / b.atBats
}

func main() {
  players := []baseballPlayer{
    baseballPlayer{name: "bob", atBats: 3.1, hits: 10},
    baseballPlayer{name: "ed", atBats: 2.1, hits: 20},
    baseballPlayer{name: "carl", atBats: 5, hits: 100},
  }

  for _, player := range players {
    println(player.name, player.average())
  }
}
