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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "bulbasaur charmander squirtle",
			expected: []string{"bulbasaur", "charmander", "squirtle"},
		},
		{
			input:    "My last TEST",
			expected: []string{"my", "last", "test"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Different lenght in sizes %v != %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Different word result %v != %v", word, expectedWord)				
			}
		}
	}
}