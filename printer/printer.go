package t9printer

import (
	"fmt"
	"github.com/nikolay-tsay/t9filter"
	"github.com/nikolay-tsay/t9reader"
)

const path = "sample.txt"

func PrintAllContacts() {
	contacts := t9reader.ReadAllContacts(path)
	if len(*contacts) == 0 {
		fmt.Println("List is empty")
		return
	}

	for name, phone := range *contacts {
		fmt.Printf("%s, %s\n", name, phone)
	}
}

func PrintFilteredContacts(searchValue int) {
	contacts := t9reader.ReadAllContacts(path)

	t9filter.Filter(contacts, searchValue)
	if len(*contacts) == 0 {
		fmt.Println("Not found")
		return
	}

	for name, phone := range *contacts {
		fmt.Printf("%s, %s\n", name, phone)
	}
}
