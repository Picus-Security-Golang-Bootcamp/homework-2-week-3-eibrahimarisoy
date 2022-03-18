# my concurrency article url
https://medium.com/@eibrahimarisoy/go-ve-concurrency-fd42e4627514
# Book Store

This repository contains Book Store application written by Go.
Book store has a lot of book. 
You can use below commands.

## Clone the project
```
$ git clone https://github.com/Picus-Security-Golang-Bootcamp/homework-1-week-2-eibrahimarisoy.git
$ cd homework-1-week-2-eibrahimarisoy
```

## Commands
### list command
```
go run main.go list
```
This command displays all books in the list.

### search command 
```
go run main.go search <bookName>
go run main.go search Lord of the Ring: The Return of 
```
This command searches for a book in the list. If the book is found, it displays the book. If the book is not found, it displays the message "We don't have that book".

### get command
```
go run main.go get <bookID>
go run main.go get 5
```
This command gets a book from the list by given ID. If the book is found, it displays the book. If the book is not found, it displays the message "We don't have that book".

### delete command
```
go run main.go delete <bookID>
go run main.go delete 5
```
This command deletes a book from the list by given ID. If the book is found, it displays the message "Book deleted successfully". If the book is not found, it displays the message "We don't have that book".

### buy command
```
go run main.go buy <bookID> <quantity>
go run main.go buy 5 2
```
This command buys a book from the list by given ID and quantity. If the book is not found, it displays the message "We don't have that book". If the quantity is not enough, it displays the message "We don't have that quantity". If the book is found and the quantity is enough, it displays the message "Thank you for your purchase". And the book quantity is decreased by the given quantity lastly print the book information.