package main
import "fmt"

func main() {

  greeting := func() { // anon function
    fmt.Println("Hello World")
  }

  greeting()
  fmt.Printf("%T\n", greeting)

  greet := makeGreeter()
  fmt.Println(greet())

  fmt.Printf("%T\n", greet)
}

func makeGreeter() func() string { // returns a anon string function
  return func() string {
    return "Hello Tipper!"
  }
}
