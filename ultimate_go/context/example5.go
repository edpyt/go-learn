package main

import (
  "context"
  "io"
  "log"
  "net/http"
  "os"
  "time"
)

func main() {
  req, err := http.NewRequest("GET", "https://www.ardanlabs.com/blog/post/index.xml", nil)
  if err != nil {
    log.Println("ERROR:",err)
    return
  }

  ctx, cancel := context.WithTimeout(req.Context(), 50 * time.Millisecond)
  defer cancel()

  req = req.WithContext(ctx)

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    log.Println("ERROR:", err)
    return
  }

  defer resp.Body.Close()

  io.Copy(os.Stdout, resp.Body)
}
