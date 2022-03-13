package book

import (
	"fmt"
	"strings"
)

type Author struct {
	id   string
	name string
}

func NewAuthor(id, name string) Author {
	return Author{id, name}
}

type Book struct {
	id         int
	name       string
	pages      string
	stockCount int
	price      float64
	stockCode  string
	ISBN       string
	Author
}

// NewBook returns a new Book instance
func NewBook(name, pages, stock_code, ISBN string, id, stock_count int, price float64, author Author) Book {
	return Book{id, name, pages, stock_count, price, stock_code, ISBN, author}
}

// bookInfo prints the book information
func (b *Book) BookInfo() {
	fmt.Printf("ID : %v \n", b.id)
	fmt.Printf("Name : %s \n", b.name)
	fmt.Printf("Pages : %s \n", b.pages)
	fmt.Printf("Stock Count : %d \n", b.stockCount)
	fmt.Printf("Price : %.2f \n", b.price)
	fmt.Printf("Stock Code : %s \n", b.stockCode)
	fmt.Printf("ISBN : %s \n", b.ISBN)
	fmt.Printf("Author ID : %s \n", b.Author.id)
	fmt.Printf("Author Name : %s \n", b.Author.name)
}

// list prints all the books
func List(books []Book) {
	for _, v := range books {
		v.BookInfo()
		fmt.Println("-", strings.Repeat("-", 50))
	}
}

// Search checks if a book is in the books slice, and returns the book
func Search(bookName string, books *[]Book) (result []Book) {
	for _, v := range *books {
		if strings.Contains(strings.ToLower(v.name), bookName) {
			result = append(result, v)
		}
	}
	return result
}

// Get returns the index of book given by the id
func Get(id int, books *[]Book) (int, error) {
	for i, v := range *books {
		if v.id == id {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Book not found")
}

// Delete deletes the book from the books slice
func Delete(books *[]Book, index int) {
	*books = append((*books)[:index], (*books)[index+1:]...)
}

// Buy decrements the stock count if the book is available
func Buy(book *Book, quantity int) error {
	if book.stockCount < quantity {
		return fmt.Errorf("Not enough stock")
	}
	book.stockCount -= quantity
	return nil
}
