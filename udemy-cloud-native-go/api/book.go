package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Book type with Title, Author and ISBN
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

// ToJSON marshalls a Book
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON unmarshalls books
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var books = map[string]Book{
	"4543456000": Book{Title: "1984", Author: "George Orwell", ISBN: "4543456000"},
	"0123456789": Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"},
}

// BooksHandleFunc http.HandleFunc for api
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// BookHandleFunc http.HandleFunc for api
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		found := DeleteBook(isbn)
		if found {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Unsupported request method."))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// AllBooks return all books
func AllBooks() []Book {
	b := make([]Book, len(books))
	i := 0
	for _, v := range books {
		b[i] = v
		i++
	}
	return b
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
}

// CreateBook saves a book into map
func CreateBook(book Book) (string, bool) {
	if _, ok := books[book.ISBN]; ok {
		return book.ISBN, false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}

// GetBook fetch from map
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

// UpdateBook updates a book in the map
func UpdateBook(book Book) bool {
	if _, ok := books[book.ISBN]; ok {
		books[book.ISBN] = book
		return true
	}
	return false
}

// DeleteBook removes a book from the map
func DeleteBook(isbn string) bool {
	if _, ok := books[isbn]; ok {
		delete(books, isbn)
		return true
	}
	return false
}
