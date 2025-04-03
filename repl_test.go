package main

import (
	"pokedexcli/internal/pokecache"
	"testing"
	"time"
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

func TestCacheData(t *testing.T) {

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "example.com",
			val: []byte("123"),
		},
		{
			key: "example2.com",
			val: []byte("more tests"),
		},
	}

	
	for _, c := range cases {
		newcache := pokecache.NewCache(2*time.Second)
		newcache.Add(c.key, c.val)
		val, ok := newcache.Get(c.key)
		if !ok {
			t.Error("Expected to find key!")
			return
		}
		if string(val) != string(c.val) {
			t.Errorf("Expected value to be the same!")
			return
		}
	}
}