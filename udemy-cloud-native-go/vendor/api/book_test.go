package api

import "testing"
import "github.com/stretchr/testify/assert"

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(
		t,
		`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`,
		string(json),
		"Error marshalling JSON",
	)
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`)
	book := FromJSON(json)
	expected := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	assert.Equal(t, expected, book)
}
