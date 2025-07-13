package main

import (
	"strings"
	"testing"
)

func TestMap_Uppercase(t *testing.T) {
	words := []string{"hello", "world"}
	expected := []string{"HELLO", "WORLD"}
	got := Map(words, func(word string) string {
		return strings.ToUpper(word)
	})

	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("Expected %s, got %s at index %d", expected[i], got[i], i)
		}
	}
}