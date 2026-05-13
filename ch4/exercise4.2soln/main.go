package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a value and one of SHA256 (which is the default), SHA384 or SHA512 to see the hash of the value or ctrl+c to exit. E.g. foobar SHA512:")
		scanner.Scan()
		line := scanner.Text()

		parts := strings.Split(line, " ")

		if len(parts) == 1 {
			fmt.Println(parts)
		}
	}
}
