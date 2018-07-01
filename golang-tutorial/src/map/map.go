package main

import "fmt"

func main() {
  m := make(map[string]int) //key: string, value: int

  m["k1"] = 7
  m["k2"] = 13
  //delete(m, "k2")
  fmt.Println("map: ", m)

  v, ok := m["k1"]
  fmt.Println("ok?:", ok, v) // ok var checks if value is there

  for key, val := range m {
    fmt.Println(key, ":", val)
  }
}
