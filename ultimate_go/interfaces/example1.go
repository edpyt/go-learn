// NOTE:
// Sample program that could benefit from polymorphioc behavior with interfaces.

package main

import "fmt"

type file struct {
  name string
}

func (file) read(b []byte) (int, error) {
  s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
  copy(b,s)
  return len(s), nil
}

type pipe struct {
  name string
}

func (pipe) read(b []byte) (int, error) {
  s := `{name: "bill", title: "developer"}`
  copy(b, s)
  return len(s), nil
}

func main() {
  f := file{"data.json"}
  p := pipe{"cfg_service"}

  retrieveFile(f)
  retrievePipe(p)
}

func retrieveFile(f file) error {
  data := make([]byte, 100)

  len, err := f.read(data)
  if err != nil {
    return err
  }

  fmt.Println(string(data[:len]))
  return nil
}

func retrievePipe(p pipe) error {
  data := make([]byte, 100)

  len, err := p.read(data)
  if err != nil {
    return err
  }

  fmt.Println(string(data[:len]))
  return nil
}

