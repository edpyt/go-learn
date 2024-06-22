package main

import "fmt"

func main() {
  scores := map[string]int{
    "anna": 21,
    "jacob": 12,
  }

  double(scores, "anna")

  fmt.Println("Score:",scores["anna"])
}

func double(scores map[string]int, player string) {
  scores[player] = scores[player] * 2
}
