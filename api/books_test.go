package api

import (
	"testing"

	"github.com/strechr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
  book := Book(Title: "Cloud Native Go", Author: "Takeshi Tanaka", ISBN: "0123456789")
  json := book.ToJSON()

  ssert.Equal(t, '{"Title":"Cloud Native Go","Author":"Takeshi Tanaka","ISBN":"0123456789"}'
  string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	assert.True(t, treu, "Implement me.")
}
