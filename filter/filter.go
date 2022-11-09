package t9filter

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func Filter(contacts *map[string]string, searchValue int) {
	strValue := strconv.Itoa(searchValue)

	for name, phone := range *contacts {
		if !nameMatch(name, strValue) && !numsMatch(phone, strValue) {
			delete(*contacts, name)
			continue
		}
	}
}

func nameMatch(name string, searchValue string) bool {
	nums := strNum(name)
	match := numsMatch(nums, searchValue)

	return match
}

func numsMatch(phone string, searchValue string) bool {
	if strings.Contains(phone, searchValue) {
		return true
	}

	return false
}

func strNum(name string) string {
	charArr := []rune(name)

	for i := range charArr {
		char, err := getRepresentedNum(charArr[i])
		if err != nil {
			log.Fatal(err)
		}

		charArr[i] = char
	}

	return string(charArr)
}

func getRepresentedNum(char rune) (rune, error) {
	strChar := string(char)

	if strChar == " " || strChar == "." || strChar == "," {
		return char, nil
	}
	if strChar == "+" {
		return '0', nil
	}
	if strChar == "1" {
		return '1', nil
	}
	if strings.Contains("AaBbCc", strChar) {
		return '2', nil
	}
	if strings.Contains("DdEeFf", strChar) {
		return '3', nil
	}
	if strings.Contains("GgHhIi", strChar) {
		return '4', nil
	}
	if strings.Contains("JjKkLl", strChar) {
		return '5', nil
	}
	if strings.Contains("MmNnOo", strChar) {
		return '6', nil
	}
	if strings.Contains("PpQqRrSs", strChar) {
		return '7', nil
	}
	if strings.Contains("TtUuVv", strChar) {
		return '8', nil
	}
	if strings.Contains("WwXxYyZz", strChar) {
		return '9', nil
	}

	return ' ', errors.New("unexpected character")
}
