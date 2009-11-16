/*
	Clone of freq.c from Kernighan and Pike: The Practice
	of Programming. I wrote two of these, this version is
	closer to the original C implementation. The good: It
	is pretty fast. The bad: It needs a huge array.
*/

package main

import "fmt"
import "os"
import "bufio"
import "unicode"

var histogram = make([]int, unicode.MaxRune+1);

func printable(char int) int {
	// kludge: no unicode.IsPrint() in library?
	if char >= 32 {
		return char;
	}
	return '-';
}

func main() {
	reader := bufio.NewReader(os.Stdin);
	for
		rune, _, error := reader.ReadRune();
		error == nil;
		rune, _, error = reader.ReadRune()
	{
		histogram[rune] = histogram[rune] + 1;
	}

	for key, value := range histogram {
		if (value > 0) {
			fmt.Printf("%.2x  %c  %d\n", key, printable(key), value);
		}
	}
}
