package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Utility
// Function that applies the given callback over the given list of strings
func Map(words []string, fn func(word string) string) []string {
	result := make([]string, len(words))
	for i, word := range words {
		result[i] = fn(word)
	}
	return result
}

// Read a new line
func ReadLine(prompt string, nl bool) string {
	reader := bufio.NewReader(os.Stdin)
	if nl {
		fmt.Print(prompt)
	} else {
		fmt.Print(prompt)
	}
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

// reverse returns a new slice with the elements of s in reverse order.
func reverse(s []string) []string {
	result := make([]string, len(s))
	for i := range s {
		result[i] = s[len(s)-1-i]
	}
	return result
}

// Task 1_1 Counter
func Task1_1() map[string]int {
	line := ReadLine("String> ", false)
	words := strings.Split(line, " ")

	low_cased := Map(words, func(word string) string {
		return strings.ToLower(strings.TrimSpace(word))
	})

	count := make(map[string]int, len(words))
	lc_cnt := make(map[string]int, len(words))

	for _, word := range low_cased {
		_, ok := lc_cnt[word]
		if ok {
			lc_cnt[word] += 1
		} else {
			lc_cnt[word] = 1
		}
	}

	for i, word := range words {
		count[word] = lc_cnt[low_cased[i]]
		fmt.Printf("%s: %d\n", word, count[word])
	}

	return  count
}

// Task 1_2 Palindrome Checker
func Task1_2() (res bool) {
	word := ReadLine("Word> ", false)

	lc := Map(strings.Split(word, ""), func(w string) string {
		return strings.ToLower(w)
	})

	res =  strings.Join(lc, "") == strings.Join(reverse(lc), "")

	if res {
		fmt.Printf("The word %s is palindrome.\n", word)
	} else {
		fmt.Printf("The word %s is not palindrome.\n", word)
	}

	return
}

func main() {
	// Task1_1()
	// Task1_2()
}
