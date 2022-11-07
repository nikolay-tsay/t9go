package main

import (
	"github.com/nikolay-tsay/t9printer"
	"os"
)

func main() {
	args := os.Args[1]
	if len(args) == 0 {
		t9printer.PrintAllContacts()
		os.Exit(0)
	}

	t9printer.PrintFilteredContacts(args)
}
