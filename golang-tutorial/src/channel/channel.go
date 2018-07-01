package main

import (
  "fmt"
  "time"
)

func main() {
  c := make(chan int)

  go func() {
    for i:= 0; i < 10; i++ {
      c <- i // passing value to the channel
    }
  }()

  go func() {
    for {
      fmt.Println(<-c)   // receiving value
    }
  }()

  time.Sleep(time.Second)
}
