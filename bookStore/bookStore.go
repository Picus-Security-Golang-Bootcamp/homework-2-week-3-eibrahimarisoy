package bookStore

import (
	"fmt"
	"strings"

	"github.com/Picus-Security-Golang-Bootcamp/homework-2-week-3-eibrahimarisoy/book"
)

type BookStore struct {
	books []*book.Book
}

// NewBookStore loads the initial books
func NewBookStore() BookStore {
	return BookStore{
		books: []*book.Book{
			book.NewBook("Book 1", "1", "1", 1, 1, 100, 1.0, false, book.NewAuthor("1", "Author 1")),
			book.NewBook("Book 2", "2", "2", 2, 2, 200, 2.0, false, book.NewAuthor("2", "Author 2")),
			book.NewBook("Book 3", "3", "3", 3, 3, 300, 3.0, false, book.NewAuthor("3", "Author 3")),
			book.NewBook("Book 4", "4", "4", 4, 4, 400, 4.0, false, book.NewAuthor("4", "Author 4")),
			book.NewBook("Book 5", "5", "5", 5, 5, 500, 5.0, false, book.NewAuthor("5", "Author 5")),
			book.NewBook("Book 6", "6", "6", 6, 6, 600, 6.0, true, book.NewAuthor("6", "Author 6")),
		},
	}
}

// List prints all the books in bookStore
func (b BookStore) List() {
	for _, v := range b.books {
		v.BookInfo()
		fmt.Println("-", strings.Repeat("-", 50))
	}
}

// Search checks if a book is in the books slice, and returns the book
func (b BookStore) Search(bookName string) (result []*book.Book) {
	for _, v := range b.books {
		if strings.Contains(strings.ToLower(v.Name), bookName) {
			result = append(result, v)
		}
	}
	return result
}

// Get returns the index of book given by the id
func (b BookStore) Get(id int) (int, error) {
	for i, v := range b.books {
		if v.ID == id {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Book not found")
}

// Delete deletes the book from the bookStore
func (b BookStore) Delete(index int) {
	instance := b.books[index]
	instance.SetIsDeleted(true)
}

// Buy decrements the stock count if the book is available
func (b BookStore) Buy(book *book.Book, quantity int) error {
	if book.IsDeleted {
		return fmt.Errorf("Book is not available")
	}
	if book.StockCount < quantity {
		return fmt.Errorf("Not enough stock")
	}
	book.SetBookStock(book.StockCount - quantity)
	return nil
}