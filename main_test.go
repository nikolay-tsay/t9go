package main

import "testing"

func Test_validInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Ok", "616", 616},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validInput(tt.input); got != tt.want {
				t.Errorf("validInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
