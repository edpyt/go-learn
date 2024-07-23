package main

import (
  "fmt"
  "runtime"
)

func main() {
  sentence := `The quick brown fox jumps over the lazy dog Stay hungry.
  Stay foolish Keep going. Be all in Boldness be my friend Screw it,
  let's do it My life is my message Leave no stone unturned Dream big.
  Pray bigger`

  print(concurrent(sentence))
}

func concurrent(text string) map[rune]int {
  m := make(map[rune]int)
  g := runtime.GOMAXPROCS(0)
  l := len(text)
  b := l / g

  ch := make(chan map[rune]int, g)

  for i := 0; i < g; i++ {
    str := i * b
    end := str + b

    if i == g - 1 {
      end = end + (l - end)
    }

    go func() {
      m := make(map[rune]int)

      defer func() {
        ch <- m
      }()

      for _, r := range text[str:end]   {
        m[r]++
      }
    }()
  }

  for i := 0; i < g; i++ {
    result := <-ch
    for rk, rv := range result {
      m[rk] = m[rk] + rv
    }
  }

  return m
}

func print(m map[rune]int) {
  var cols int

  for r := 65; r< 65+26;r++ {
    v := m[rune(r)]
    fmt.Printf("%q:%d", rune(r), v)

    v = m[rune(r+32)]
    fmt.Printf("%q:%d, ", rune(r + 32), v)

    cols++
    if cols == 5 {
    fmt.Print("\n")}
    cols = 0
  }
}
