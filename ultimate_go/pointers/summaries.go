type user struct {
  name  string
  email string
}

func stayOnStack() user {
  // Функция stayOnStack использует семантику значения
  // для возврата значения user вызывающему.
  // Другими словами, вызывающий получает собственную копию
  // значения user, созданного функцией stayOnStack

  u := user{
    name:   "Bill",
    email:  "bill@email.com",
  }

  return u
}

func escapeToHeap() *user {
  u := user{
    name:   "Bill",
    email:  "bill@email.com",
  }

  return &u
}
