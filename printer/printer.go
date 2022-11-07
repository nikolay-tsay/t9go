package t9printer

import (
	"fmt"
	"github.com/nikolay-tsay/t9filter"
	"github.com/nikolay-tsay/t9reader"
)

func PrintAllContacts() {
	contacts := t9reader.ReadAllContacts()

	for name, phone := range contacts {
		fmt.Printf("%s, %s\n", name, phone)
	}
}

func PrintFilteredContacts(searchValue string) {
	contacts := t9reader.ReadAllContacts()

	t9filter.Filter(contacts, searchValue)

	for name, phone := range contacts {
		fmt.Printf("%s, %s\n", name, phone)
	}
}
