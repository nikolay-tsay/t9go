package t9reader

import (
	"bufio"
	"log"
	"os"
)

type Contact struct {
	Name        string
	PhoneNumber string
}

func ReadAllContacts() []Contact {
	lines := readFile()

	i := 0
	j := 0
	var contacts []Contact
	for range lines {
		contact := convertContact(lines[i], lines[i+1])
		contacts[j] = *contact

		i += 2
		j++
	}

	return contacts
}

func convertContact(name string, phoneNumber string) *Contact {
	var contact Contact
	contact.Name = name
	contact.PhoneNumber = phoneNumber

	return &contact
}

func readFile() []string {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
