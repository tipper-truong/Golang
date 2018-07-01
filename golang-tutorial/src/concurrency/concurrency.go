package main

import (
  "fmt"
  "sync"
  "time"
)

var wg sync.WaitGroup

func main() {
  wg.Add(2) // add two to the wait group and run foo and bar
  go foo() // creates a goroutine or lightweight thread when running go foo()
  go bar() // creates a goroutine or lightweight thread when running go foo()
  wg.Wait() // semaphore of waitgroup becomes 0 b/c foo() and bar() already ran
}

func foo() {
  for i := 0; i < 45; i++ {
    fmt.Println("Foo:", i)
    time.Sleep(time.Duration(3 * time.Millisecond))
  }
  wg.Done() // tells the WaitGroup with foo() is done, removes foo()
}

func bar() {
  for i:= 0; i < 45; i++ {
    fmt.Println("Bar:", i)
    time.Sleep(time.Duration(3 * time.Millisecond))
  }
  wg.Done() // tells the WaitGroup with bar() is done, removes bar()
}
