package main

import (
	"fmt"

	"github.com/zmflexopack/the-deeper-love-of-go/books"
)

func main() {
    book := books.Book{
        Title:  "Engineering in Plain Sight",
        Author: "Grady Hillhouse",
        Copies: 2,
    }
    fmt.Println(books.BookToString(book))
}