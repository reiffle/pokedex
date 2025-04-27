package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		}, {
			input:    "HeLlo woRld ",
			expected: []string{"hello", "world"},
		}, {
			input:    "           HELLO             world      ",
			expected: []string{"hello", "world"},
		}, {
			input:    "",
			expected: []string{},
		}, {
			input:    "I AM          ScuZZlebutt",
			expected: []string{"i", "am", "scuzzlebutt"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length does not match: Got %v, Want %v", len(actual), len(c.expected))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word mismatch. Expected %v, Got %v", expectedWord, word)
				return
			}
		}
	}
}
