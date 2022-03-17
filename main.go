package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Picus-Security-Golang-Bootcamp/homework-2-week-3-eibrahimarisoy/bookStore"
)

// define usage information
var usage = `Usage: ./ [commands...] [parameters...]

Commands:
	-list
	-search <bookName>
	-get <bookID>
	-delete <bookID>
	-buy <bookID> <quantity>

Parameters:
	-keyword: string
	-bookID: int
	-quantity: int
`

func main() {
	// define usage
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	// load books

	args := os.Args[1:]

	if err := bookStore.Run(bookStore.NewBookStore(), args); err != nil {
		usageAndExit(err.Error())
	}
}

func usageAndExit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	fmt.Fprintf(os.Stderr, "\n\n")
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
