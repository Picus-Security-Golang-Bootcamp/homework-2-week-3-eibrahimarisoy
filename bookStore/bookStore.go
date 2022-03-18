package bookStore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Picus-Security-Golang-Bootcamp/homework-2-week-3-eibrahimarisoy/book"
)

type BookStore struct {
	Books []*book.Book
}

// NewBookStore loads the bookStore from the file using the map[string] and
// also use constructor func for book and author
func NewBookStore() (*BookStore, error) {
	bs := BookStore{}

	data := []map[string]interface{}{}
	contents, err := ioutil.ReadFile("books.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(contents, &data); err != nil {
		return nil, err
	}

	for _, v := range data {
		bs.Books = append(bs.Books, book.NewBook(v))
	}

	return &bs, nil
}

// NewBookStore loads the bookStore from the file using struct
// func NewBookStore() (*BookStore, error) {
// 	bs := BookStore{}
// 	jsonFile, err := os.Open("books.json")

// 	if err != nil {
// 		return nil, err
// 	}

// 	data, err := ioutil.ReadAll(jsonFile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := json.Unmarshal(data, &bs); err != nil {
// 		return nil, err
// 	}

// 	return &bs, nil
// }

// Close saves the bookStore to the file
func (b BookStore) WriteFile() {
	file, _ := json.MarshalIndent(b.Books, "", " ")
	_ = ioutil.WriteFile("books.json", file, 0644)

}

// List prints all the books in bookStore
func (b BookStore) List() {
	for _, v := range b.Books {
		v.BookInfo()
		fmt.Println("-", strings.Repeat("-", 50))
	}
}

// Search checks if a book is in the books slice, and returns the book
func (b BookStore) Search(bookName string) (result []*book.Book) {
	for _, v := range b.Books {
		if strings.Contains(strings.ToLower(v.Name), bookName) {
			result = append(result, v)
		}
	}
	return result
}

// Get returns the index of book given by the id
func (b BookStore) Get(id int) (int, error) {
	for i, v := range b.Books {
		if v.ID == id {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Book not found")
}

// Delete deletes the book from the bookStore
func (b BookStore) Delete(index int) {
	instance := b.Books[index]
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
