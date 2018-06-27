package main

import "fmt"

func hello() {
  fmt.Print("hello")
}

func world() {
  fmt.Println("world")
}

func main() {
  defer world() // when the main function exits, it closes the file
  hello()
}
