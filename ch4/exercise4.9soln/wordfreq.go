package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	var inputSource io.Reader = os.Stdin
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		inputSource = file
	}

	wordFreq := make(map[string]int)

	scanner := bufio.NewScanner(inputSource)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		word = strings.Trim(word, "!.?(),;")
		wordFreq[word]++
	}

	//scanner.Scan can return false when either there's no more to read or error happened when it tried to scan
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error when scanning: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%20v\t%v\n", "Word", "Frequency")
	for word, count := range wordFreq {
		fmt.Printf("%20v\t%v\n", word, count)
	}

}
