package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(isAnagram("john", "john"))
	fmt.Println(isAnagram("listen", "silent"))
	fmt.Println(isAnagram("aab", "abb"))

}

func isAnagram(a, b string) bool {

	if a == b || len(a) != len(b) {
		return false
	}

	for _, v := range b {
		index := strings.IndexRune(a, v)
		if index == -1 {
			return false
		}
		a = a[:index] + a[index+utf8.RuneLen(v):]

	}
	return true
}
