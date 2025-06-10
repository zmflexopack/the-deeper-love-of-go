package main

import (
	"fmt"

	"github.com/zmflexopack/the-deeper-love-of-go/books"
)

func main() {
	catalog := books.GetCatalog()
	for _, book := range books.GetAllBooks(catalog) {
	fmt.Println(book)
	}
}