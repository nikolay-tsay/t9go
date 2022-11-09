package t9reader

import (
	"reflect"
	"testing"
)

func TestReadAllContacts(t *testing.T) {
	tests := []struct {
		name string
		want *map[string]string
	}{
		{"AllRead", &map[string]string{"Petr Dvorak": "603123456", "Jana Novotna": "777987654", "Bedrich Smetana ml.": "541141120"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadAllContacts("C:\\Users\\Meta IT\\Desktop\\Projects\\Go\\t9go\\sample.txt"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAllContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"AllRead", []string{"Petr Dvorak", "603123456", "Jana Novotna", "777987654", "Bedrich Smetana ml.", "541141120"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFile("C:\\Users\\Meta IT\\Desktop\\Projects\\Go\\t9go\\sample.txt"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
