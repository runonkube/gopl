package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

	names := []byte("ab  c dan  simon      aaron luke  Shipp John")

	fmt.Println(string(removeAdjacentUnicodeSpaces(names)))

	letters := []byte("  abc  ")
	fmt.Printf("%q\n", string(removeAdjacentUnicodeSpaces(letters)))

	letters = []byte("  foo   bar    ")
	fmt.Printf("%q\n", string(removeAdjacentUnicodeSpaces(letters)))

	letters = []byte("a  \tb")
	fmt.Printf("%q\n", string(removeAdjacentUnicodeSpaces(letters)))

}

func removeAdjacentUnicodeSpaces(theBytes []byte) []byte {
	write := 0
	for read := 0; read < len(theBytes); {
		ch, size := utf8.DecodeRune(theBytes[read:])
		if unicode.IsSpace(ch) {
			theBytes[write] = ' '
			write++
			for read < len(theBytes) {
				next, nextSize := utf8.DecodeRune(theBytes[read:])
				if !unicode.IsSpace(next) {
					break
				}
				read += nextSize
			}
		} else {
			copy(theBytes[write:], theBytes[read:read+size])
			write += size
			read += size
		}
	}
	return theBytes[:write]
}
