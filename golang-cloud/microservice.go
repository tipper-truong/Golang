package main

import (
  "os"
  "fmt"
  "net/http"
  "github.com/tqtruong95/Golang/golang-cloud/api"
)

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("api/echo", echo)
  http.HandleFunc("/api/books", api.BooksHandleFunc)
  http.HandleFunc("/api/books/", api.BooksHandleFunc)
  http.ListenAndServe(port(), nil)
}

func port() string {
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "8080"
  }
  return ":" + port
}

// ResponseWriter, respond content from our client
// httpRequest, get,post,delete,put
func index(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Hello Cloud Native Go")
}

func echo(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Query()["message"][0]

  w.Header().Add("Content-Type", "text/plain")
  fmt.Fprintf(w, message)
}
