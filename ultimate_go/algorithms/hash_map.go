package main

import (
  "fmt"
  "hash/maphash"
)

func main() {
  h := New()

  k1, v1 := "key1", 1
  k2, v2 := "key2", 2
  h.Store(k1, v1)
  h.Store(k2, v2)

  v, err := h.Retrieve(k1)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("key:", k1, "value:", v)

  v1b := 11
  h.Store(k1, v1b)

  v, err = h.Retrieve(k1)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("key:", k1, "value:", v)

  if err := h.Delete(k1); err != nil {
    fmt.Println(err)
    return
  }

  v, err = h.Retrieve(k1)
  if err != nil {
    fmt.Println(err)
  }

  fn := func(key string, value int) bool {
    fmt.Println("key:", key, "value:", value)
    return true
  }
  h.Do(fn)
}

const numBuckets = 256

type entry struct {
  key string
  value int
}

type Hash struct {
  buckets [][]entry
  hash    maphash.Hash
}

func New() *Hash {
  return &Hash{
    buckets: make([][]entry, numBuckets),
  }
}

func (h *Hash) Store(key string, value int) {
  idx := h.hashKey(key)

  bucket := h.buckets[idx]

  for idx := range bucket {
    if bucket[idx].key == key {
      bucket[idx].value = value
      return
    }
  }

  h.buckets[idx] = append(bucket, entry{key, value})
}

func (h *Hash) Retrieve(key string) (int, error) {
  idx := h.hashKey(key)

  for _, entry := range h.buckets[idx] {
    if entry.key == key {
      return entry.value, nil
    }
  }

  return 0, fmt.Errorf("%q not found", key)
}

func (h *Hash) Delete(key string) error {
  bucketIdx := h.hashKey(key)

  bucket := h.buckets[bucketIdx]

  for entryIdx, entry := range bucket {
    if entry.key == key {
      bucket = removeEntry(bucket, entryIdx)
      h.buckets[bucketIdx] = bucket
      return nil
    }
  }

  return fmt.Errorf("%q not found", key)
}

func (h *Hash) Len() int {
  sum := 0
  for _, bucket := range h.buckets {
    sum += len(bucket)
  }
  return sum
}

func (h *Hash) Do(fn func(key string, value int) bool) {
  for _, bucket := range h.buckets {
    for _, entry := range bucket {
      if ok := fn(entry.key, entry.value); !ok {
        return
      }
    }
  }
}

func (h *Hash) hashKey(key string) int {
  h.hash.Reset()
  h.hash.WriteString(key)
  
  n := h.hash.Sum64()

  return int(n % numBuckets)
}

func removeEntry(bucket []entry, idx int) []entry {
  copy(bucket[idx:], bucket[idx+1:])
  bucket = bucket[:len(bucket) - 1]
  return reduceAllocation(bucket)
}

func reduceAllocation(bucket []entry) []entry {
  if cap(bucket) < 2 * len(bucket) {
    return bucket
  }
  newBucket := make([]entry, len(bucket))
  copy(newBucket, bucket)
  return newBucket
}
