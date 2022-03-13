package bookStore

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Bootcamp/homework-2-week-3-eibrahimarisoy/book"
)

// Run runs the bookStore given the command and the arguments
func Run(books *[]book.Book, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("No command provided")
	}

	switch args[0] {

	case "list":
		// list all the books
		book.List(*books)

	case "search":
		// if the user has not provided <bookName>
		if len(args) < 2 {
			return fmt.Errorf("No book name provided")
		}
		results := book.Search(strings.Join(args[1:], " "), books)
		if len(results) == 0 {
			return fmt.Errorf("No book found")
		}
		book.List(results)

	case "get":
		// if the user has not provided <bookID>
		if len(args) < 2 {
			return fmt.Errorf("No book id provided")
		}

		bookId, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		index, err := book.Get(bookId, books)
		if err != nil {
			return err
		}
		(*books)[index].BookInfo()

	case "delete":
		// if the user has not provided <bookID>
		if len(args) < 2 {
			return fmt.Errorf("No book id provided")
		}

		bookId, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		index, err := book.Get(bookId, books)
		if err != nil {
			return err
		}
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("Deleting book:")
		(*books)[index].BookInfo()
		fmt.Println(strings.Repeat("-", 50))

		book.Delete(books, index)

		fmt.Println("Remaining books:")
		fmt.Println()
		book.List(*books)

	case "buy":
		// if the user has not provided <bookID> and <quantity>
		if len(args) < 3 {
			return fmt.Errorf("No book id or quantity provided")
		}

		bookId, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		quantity, err := strconv.Atoi(args[2])
		if err != nil {
			return err
		}

		index, err := book.Get(bookId, books)

		if err != nil {
			return err
		}

		instance := (*books)[index]
		if err := book.Buy(&instance, quantity); err != nil {
			return fmt.Errorf("Not enough stock")
		}

		instance.BookInfo()

	}
	return nil
}
