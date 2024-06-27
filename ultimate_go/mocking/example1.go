package main

// pubsub package ===============================================
type PubSub struct {
  host string
}

func New(host string) *PubSub {
  return &PubSub{
    host: host,
  }
}

func (ps *PubSub) Publish(key string, v interface{}) error {
  return nil
}

func (ps *PubSub) Subscribe(key string) error {
  return nil
}
// ==============================================================

type publisher interface {
  Publish(key string, v interface{}) error
  Subscribe(key string) error
}

type mock struct{}

func (m *mock) Publish(key string, v interface{}) error {
  return nil
}

func (m *mock) Subscribe(key string) error {
  return nil
}

func main() {
  pubs := []publisher{
    New("localhost"),
    &mock{},
  }

  for _, p := range pubs {
    p.Publish("key", "value")
    p.Subscribe("key")
  }
}
