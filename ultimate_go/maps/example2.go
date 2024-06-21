package main

func main() {
  scores := make(map[string]int)

  score := scores["anna"]

  println("Score:",score)

  score, ok := scores["anna"]

  println("Score:", score, "Present:", ok)

  // Можно инкрементировать нулевое значение в карте
  scores["anna"]++

  // Без этого поведения нам бы пришлось писать код
  // типа такого:
  if n, ok := scores["anna"]; ok {
    scores["anna"] = n + 1
  } else {
    scores["anna"] = 1
  }

  score, ok = scores["anna"]
  println("Score:", score, "Present:", ok)
}
