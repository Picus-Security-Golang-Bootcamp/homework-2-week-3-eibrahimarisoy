package main

import (
	"fmt"
	"os"

	"github.com/Picus-Security-Golang-Bootcamp/homework-2-week-3-eibrahimarisoy/book"
	"github.com/Picus-Security-Golang-Bootcamp/homework-2-week-3-eibrahimarisoy/bookStore"
)

// define console information
var usage = `Usage: ./ [commands...] [parameters...]

Commands:
	-list
	-search <bookName>
	-get <bookID>
	-delete <bookID>
	-buy <bookID> <quantity>

Parameters:
	- bookName: string
	- bookID: int
	- quantity: int
`

func main() {
	// define books
	books := []book.Book{
		book.NewBook("Book 1", "100", "1", "1", 1, 1, 1.0, book.NewAuthor("1", "Author 1")),
		book.NewBook("Book 2", "200", "2", "2", 2, 2, 2.0, book.NewAuthor("2", "Author 2")),
		book.NewBook("Book 3", "300", "3", "3", 3, 3, 3.0, book.NewAuthor("3", "Author 3")),
		book.NewBook("Book 4", "400", "4", "4", 4, 4, 4.0, book.NewAuthor("4", "Author 4")),
		book.NewBook("Book 5", "500", "5", "5", 5, 5, 5.0, book.NewAuthor("5", "Author 5")),
		book.NewBook("Book 6", "600", "6", "6", 6, 6, 6.0, book.NewAuthor("6", "Author 6")),
	}

	args := os.Args[1:]

	if err := bookStore.Run(&books, args); err != nil {
		usageAndExit(err.Error())
	}

}

func usageAndExit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	fmt.Fprintf(os.Stderr, "\n\n")
	fmt.Fprintf(os.Stderr, fmt.Sprintf(usage))
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
