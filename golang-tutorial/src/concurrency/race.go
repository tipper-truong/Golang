package main

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func main() {
  wg.Add(2)
  go incrementor("Foo:")
  go incrementor("Bar:")
  wg.Wait()
  fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
  for i:= 0; i < 20; i++ {
    time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond) // helps execute differently
    mutex.Lock()
    counter++
    fmt.Println(s, i, "Counter:", counter)
    mutex.Unlock()
  }
  wg.Done()
}
