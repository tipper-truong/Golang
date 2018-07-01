package main

import (
  "encoding/json"
  "fmt"
)

type Person struct {
  First string
  Last string
  Age int
  notExported int
}

func main() {
  p1 := Person{"James", "Bond", 20, 007}
  byteSlice, _ := json.Marshal(p1)
  fmt.Println(byteSlice)
  fmt.Printf("%T \n", byteSlice) // byte = 8 bit
  fmt.Println(string(byteSlice)) // doesn't export notExported b/c it's lowercase var
}
