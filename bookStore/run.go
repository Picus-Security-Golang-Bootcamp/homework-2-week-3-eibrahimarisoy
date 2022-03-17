package bookStore

import (
	"fmt"
	"strconv"
	"strings"
)

// Run runs the bookStore given the command and the arguments
func Run(bs BookStore, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("No command provided")
	}

	switch args[0] {

	case "list":
		// list all the books
		bs.List()

	case "search":
		// if the user has not provided <bookName>
		if len(args) < 2 {
			return fmt.Errorf("No book name provided")
		}

		results := bs.Search(strings.Join(args[1:], " "))

		if len(results) == 0 {
			return fmt.Errorf("No book found")
		}

		for _, book := range results {
			book.BookInfo()
			fmt.Println("-", strings.Repeat("-", 50))
		}

	case "get":
		// if the user has not provided <bookID>
		if len(args) < 2 {
			return fmt.Errorf("No book id provided")
		}

		bookId, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		index, err := bs.Get(bookId)
		if err != nil {
			return err
		}
		bs.books[index].BookInfo()

	case "delete":
		// if the user has not provided <bookID>
		if len(args) < 2 {
			return fmt.Errorf("No book id provided")
		}

		bookId, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		index, err := bs.Get(bookId)
		if err != nil {
			return err
		}
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("Deleting book:")

		bs.books[index].BookInfo()
		fmt.Println(strings.Repeat("-", 50))

		bs.Delete(index)

		fmt.Println("Remaining books:")
		fmt.Println()
		bs.List()

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

		index, err := bs.Get(bookId)

		if err != nil {
			return err
		}

		instance := bs.books[index]
		if err := bs.Buy(instance, quantity); err != nil {
			return err
		}

		instance.BookInfo()

		// default:
		// 	return fmt.Errorf("Invalid command")
	}
	return nil
}
