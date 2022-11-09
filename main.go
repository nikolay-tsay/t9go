package main

import (
	"github.com/nikolay-tsay/t9printer"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		t9printer.PrintAllContacts()
		os.Exit(0)
	}

	numValue := validInput(os.Args[1])

	t9printer.PrintFilteredContacts(numValue)
}

func validInput(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("'%s' is not valid input. Use digits only\n", input)
	}

	return num
}
