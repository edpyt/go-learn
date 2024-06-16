// IOTA
package main

import "fmt"

func main() {
  const (
    A1 = iota  // 0
    B1 = iota  // 1
    C1 = iota  // 2
  )

  fmt.Println("1:", A1, B1, C1)

  const (
    A2 = iota  // 0
    B2         // 1
    C2         // 2
  )

  fmt.Println("2:", A2, B2, C2)

  // Нет необходимости повторять ключевое слово iota.
  // Последовательный характер целых констант подразумевается
  // при его первом применении.
  const (
    A3 = iota + 1  // 1
    B3             // 2
    C3             // 3
  )

  fmt.Println("3:", A3, B3, C3)

  // Если вы не хотите использовать математический шаблон,
  // вы можете выполнить математические вычисления и применить
  // их в дальнейших увеличениях значения iota.
  const (
    Ldate = 1 << iota  // 1: Сдвиг 1 налево на 0.
    Ltime              // 2
    Lmicroseconds      // 4
    Llongfile          // 8
    Lshortfile         // 16
    LUTC               // 32
  )

  fmt.Println(Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
