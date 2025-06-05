package main

import (
	"fmt"
	"os"

	"github.com/zmflexopack/the-deeper-love-of-go/books"
)

func main () {
	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}
	ID := os.Args[1]
	book, ok := books.GetBook(ID)
	if !ok {
		fmt.Println("Sorry,I couldn't find that book in the Catalog")
		return
	}
	fmt.Println(books.BookToString(book))
}