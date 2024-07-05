package main

import (
  "context"
  "fmt"
) 

type TraceID string
type TraceIDKey int

func main() {
  traceID := TraceID("f47ac10b-58cc-0372-8567-0e02b2c3d479")

  const traceIDKey TraceIDKey = 0

  ctx := context.WithValue(context.Background(), traceIDKey, traceID)

  if uuid, ok := ctx.Value(traceIDKey).(TraceID); ok {
    fmt.Println("TraceID:", uuid)
  }

  if _, ok := ctx.Value(0).(TraceID); !ok {
    fmt.Println("TraceID Not Found")
  }
}
