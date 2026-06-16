// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	var digits, puncts, symbols, letters, spaces, controls int

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // return rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsDigit(r):
			digits++
		case unicode.IsPunct(r):
			puncts++
		case unicode.IsSpace(r):
			spaces++
		case unicode.IsSymbol(r):
			symbols++
		case unicode.IsControl(r):
			controls++
		}

	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\n%-12v\tCount\n", "Categories")
	fmt.Printf("%-12v\t%d\n", "Letters", letters)
	fmt.Printf("%-12v\t%d\n", "Digits", digits)
	fmt.Printf("%-12v\t%d\n", "Spaces", spaces)
	fmt.Printf("%-12v\t%d\n", "Symbols", symbols)
	fmt.Printf("%-12v\t%d\n", "Punctuations", puncts)
	fmt.Printf("%-12v\t%d\n", "Controls", controls)

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
