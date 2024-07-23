package main

import (
  "errors"
  "fmt"
  "math/rand"
  "sync"
  "sync/atomic"
  "time"
)

func main() {
  const maxChairs = 3

  shop := OpenShop(maxChairs)
  defer shop.Close()

  t := time.NewTimer(50 * time.Millisecond)
  <- t.C
}

var (
  ErrShopClosed = errors.New("shop closed")
  ErrNoChair = errors.New("no chair available")
)

type customer struct {
  name string
}

type Shop struct {
  open    int32
  chairs  chan customer
  wgClose sync.WaitGroup
}

func OpenShop(maxChairs int) *Shop {
  fmt.Println("Opening the shop")

  s := Shop{
    chairs: make(chan customer, maxChairs),
  }
  atomic.StoreInt32(&s.open, 1)

  s.wgClose.Add(1)
  go func() {
    defer s.wgClose.Done()

    fmt.Println("Barber ready to work")

    for cust := range s.chairs {
      s.serviceCustomer(cust)
    }
  }()

  go func() {
    var id int64

    for {
      time.Sleep(time.Duration(rand.Intn(75)) * time.Millisecond)

      name := fmt.Sprintf("cust-%d", atomic.AddInt64(&id, 1))

      if err := s.newCustomer(name); err != nil {
        if err == ErrShopClosed {
          break
        }
      }
    }
  }()

  return &s
}

func (s *Shop) Close() {
  fmt.Println("Closing the shop")
  defer fmt.Println("Shop closed")

  atomic.StoreInt32(&s.open, 0)

  close(s.chairs)
  s.wgClose.Wait()
}

func (s *Shop) serviceCustomer(cust customer) {
  fmt.Printf("Barber servicing customer %q\n", cust.name)

  time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

  fmt.Printf("Barber finished customer %q\n", cust.name)

  if len(chairs) == 0 && atomic.LoadInt32(&s.open) == 1 {
    fmt.Println("Barber taking a nap")
  }
}

func (s *Shop) newCustomer(name string) error {
  if atomic.LoadInt32(&s.open) == 0 {
    fmt.Printf("Customer %q leaves, shop closed\n", name)
    return ErrShopClosed
  }

  fmt.Printf("Customer %q entered shop\n", name)

  select {
  case s.chairs <- customer{name: name}:
    fmt.Printf("Customer %q takes a seat and waits\n", name)
  default:
    fmt.Printf("Customer %q leaves, no seat\n", name)
  }
  
  return nil
}
