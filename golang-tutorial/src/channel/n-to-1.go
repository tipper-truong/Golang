package main

import (
  "fmt"
  "sync"
)

// Many functions in different goroutines writing to the same channel
func main() {

  c := make(chan int)

  var wg sync.WaitGroup
  wg.Add(2)

  go func() {
    for i:= 0; i < 10; i++ {
      c <- i
    }
    wg.Done()
  }()

  go func() {
    for i:= 0; i < 10; i++ {
      c <- i
    }
    wg.Done()
  }()

  go func() {
    wg.Wait()
    close(c)
  }()

  for n := range c {
    fmt.Println(n)
  }
}
