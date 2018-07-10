package api

import (
  "encoding/json"
  "net/http"

)

// Book type with Name, Author, and ISBN
type Book struct {
  Title string `json:"title"` // marshall at lowercase
  Author string `json:"author"` // marshal at lowercase
  ISBN string `json:"isbn"` // marshal at lowercase
}

// ToJSON to be used for marshalling of Book type
// Marshalling means converting to byte stream
func (b Book) ToJSON() []byte {
  ToJSON, err := json.Marshal(b)
  if err != nil {
    panic(err)
  }
  return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
// Unmarshalling means decrypting byte stream from marshalling
func FromJSON(data []byte) Book {
  book := Book{}
  err := json.Unmarshal(data, &book)
  if err != nil {
    panic(err)
  }
  return book
}

var Books = []Book{
  Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adam", ISBN: "0345391802"},
  Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000000"},
}
// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
  book, err := json.Marshal(Books)
  if err != nil {
    panic(err)
  }
  w.Header().Add("Content-Type", "application/json; charset=utf-8")
  w.Write(book)
}
