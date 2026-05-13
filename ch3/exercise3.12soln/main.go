// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf(" %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	n := len((s))

	commaPos := n % 3
	if commaPos == 0 {
		commaPos = 3
	}

	for count, v := range s {

		if count == commaPos {
			buf.WriteByte(',')
			commaPos += 3
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
