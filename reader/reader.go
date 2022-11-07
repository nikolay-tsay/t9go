package t9reader

import (
	"bufio"
	"log"
	"os"
)

func ReadAllContacts() map[string]string {
	lines := readFile()

	i := 0
	j := 0
	contacts := make(map[string]string)
	for range lines {
		contacts[lines[i]] = lines[i+1]

		i += 2
		j++
	}

	return contacts
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
