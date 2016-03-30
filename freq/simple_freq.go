/*
	Clone of freq.c from Kernighan and Pike: The Practice
	of Programming. I wrote two of these, this version is
	closer to the original C implementation. The good: It
	is pretty fast. The bad: It needs a huge array.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var histogram = make([]int, unicode.MaxRune+1)

func printable(char int) int {
	if unicode.IsPrint(rune(char)) {
		return char
	}
	return '-'
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for c, _, err := r.ReadRune(); err == nil; c, _, err = r.ReadRune() {
		histogram[c]++
	}

	for key, value := range histogram {
		if value > 0 {
			fmt.Printf("%.2x  %c  %d\n", key, printable(key), value)
		}
	}
}
