package main

import (
	"pokedexcli/internal/pokecache"
	"testing"
	"time"
)

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

func TestReapLoop(t *testing.T) {
	
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
		_, ok := newcache.Get(c.key)
		if !ok {
			t.Error("Expected to find key!")
			return
		}
		
		time.Sleep(3*time.Second)

		_, ok = newcache.Get(c.key)

		if ok {
			t.Errorf("Expected not to find key!")
			return
		}
	}
}