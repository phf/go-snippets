/*
	Clone of freq.c from Kernighan and Pike: The Practice
	of Programming. I wrote two of these, this version is
	closer to a Python version. The good: Only characters
	actually occurring use memory. The bad: Pretty slow.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

var histogram = make(map[int]int)

func printable(char int) int {
	if unicode.IsPrint(rune(char)) {
		return char
	}
	return '-'
}

func build_histogram() {
	r := bufio.NewReader(os.Stdin)
	for c, _, err := r.ReadRune(); err == nil; c, _, err = r.ReadRune() {
		histogram[int(c)]++
	}
}

func sort_keys(dict map[int]int) []int {
	keys := make([]int, len(dict))
	i := 0
	for key, _ := range dict {
		keys[i] = key
		i++
	}
	sort.Ints(keys)
	return keys
}

func print_histogram() {
	// there must be a more elegant idiom for this?
	sorted_keys := sort_keys(histogram)

	for _, key := range sorted_keys {
		fmt.Printf("%.2x  %c  %d\n", key, printable(key), histogram[key])
	}
}

func main() {
	build_histogram()
	print_histogram()
}
