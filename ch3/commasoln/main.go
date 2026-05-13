// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf(" %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)

	//Check if has a sign
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteString(s[0:1])
		s = s[1:]
		n -= 1
	}

	//check if its a floating point number
	decimalPointPos := strings.Index(s, ".")
	if decimalPointPos != -1 {
		n = len(s[:decimalPointPos])
	}

	nextCommaPos := n % 3
	if nextCommaPos == 0 {
		nextCommaPos = 3
	}

	for count, v := range s {
		if count == nextCommaPos && count != decimalPointPos {
			buf.WriteByte(',')
			nextCommaPos += 3
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
