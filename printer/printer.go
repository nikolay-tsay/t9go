package t9printer

import "github.com/nikolay-tsay/t9reader"

func PrintAllContacts() {
	contacts := t9reader.ReadAllContacts()
}

func PrintFilteredContacts(searchValue string) {

}
