package books

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
}

func BookToString(book Book) string {
	return fmt.Sprintf("%s by %s - %d copies", 
		book.Title, book.Author, book.Copies)
}