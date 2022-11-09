package t9reader

import (
	"bufio"
	"log"
	"os"
)

func ReadAllContacts(path string) *map[string]string {
	lines := readFile(path)

	i := 0
	contacts := make(map[string]string)
	for range lines {
		contacts[lines[i]] = lines[i+1]

		if i+2 >= len(lines) {
			break
		}

		i += 2
	}

	return &contacts
}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
