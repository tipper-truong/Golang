package api

import (
  "encoding/json"
  "net/http"
  "io/ioutil"
  "fmt"
)

// Book type with Name, Author, and ISBN
type Book struct {
  Title string `json:"title"` // marshall at lowercase
  Author string `json:"author"` // marshal at lowercase
  ISBN string `json:"isbn"` // marshal at lowercase
}

var books = map[string]Book{
  "0345391802": Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adam", ISBN: "0345391802"},
  "0000000000000" :Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000000"},
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

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
  switch method := r.Method; method {
    case http.MethodGet: // GET
      books := AllBooks()
      writeJSON(w, books)

    case http.MethodPost:
      body, err := ioutil.ReadAll(r.Body) // get user input for book info
      fmt.Println("Body: %s", body)
      if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
      }
      book := FromJSON(body)
      isbn, created := CreateBook(book)
      if created {
        w.Header().Add("Location", "/api/books"+isbn)
        w.WriteHeader(http.StatusCreated)
      } else {
        w.WriteHeader(http.StatusConflict)
      }
  default:
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Unsupported request method"))
  }
}

func AllBooks() []Book {
  values := make([]Book, len(books))
  i := 0
  for _, book := range books {
    values[i] = book
    i++
  }
  return values
}

func CreateBook(book Book) (string, bool) {
  _, exists := books[book.ISBN]
  if exists {
    return "", false
  }
  books[book.ISBN] = book
  return book.ISBN, true
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
