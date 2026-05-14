package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	hashType := flag.String("h", "sha256", "hash type: sha256, sha384 sha512")
	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintln(os.Stderr, "Unrecognised positional args.", flag.Args())
		os.Exit(1)
	}

	fmt.Print("Enter input, press ctrl+d when done:")
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Cannot read data from std in")
		return
	}
	trimmedBytes := bytes.TrimRight(data, "\n")
	fmt.Println()

	switch strings.ToLower(*hashType) {
	case "sha512":
		fmt.Printf("sha512 of '%s' is %x\n", trimmedBytes, sha512.Sum512(trimmedBytes))
	case "sha384":
		fmt.Printf("sha384 of '%s' is %x\n", trimmedBytes, sha512.Sum384(trimmedBytes))
	default:
		fmt.Printf("sha256 of '%s' is %x\n", trimmedBytes, sha256.Sum256(trimmedBytes))
	}
}
