package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  c := fanIn(boring("Joe"), boring("Ann")) // two separate channels
  for i := 0; i < 10; i++ {
    fmt.Println(<-c)
  }
  fmt.Println("You're both boring; I'm leaving")
}

func boring(msg string) <-chan string {
  c := make(chan string)
  go func() {
    for i := 0; ; i++ {
      c <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
  }()
  return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
  c := make(chan string)
  go func() {
    for {
      c <- <-input1 // <-input1 is a value and puts it on channel c
    }
  }()

  go func() {
    for {
      c <- <-input2 // <-input2 is a value and puts it on channel c
    }
  }()
  return c
}
