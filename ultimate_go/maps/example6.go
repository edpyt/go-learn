package main

type player struct {
  name string
  score int
}

func main() {
  players := map[string]player{
    "anna":   {"Anna", 42},
    "jacob":  {"Jacob", 21},
  }

  anna := &players["anna"]
  anna.score++

  player := players["anna"]
  player.score++
  players["anna"] = player
}
