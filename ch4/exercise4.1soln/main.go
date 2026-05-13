package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	hash1 := sha256.Sum256([]byte("foo"))
	hash2 := sha256.Sum256([]byte("foo"))
	fmt.Printf("%x\n%x\nNumber of different bits=%d\n", hash1, hash2, countDifferentBits(hash1, hash2))

	fmt.Println()
	hash1 = sha256.Sum256([]byte("foo"))
	hash2 = sha256.Sum256([]byte("Foo"))
	fmt.Printf("%x\n%x\nNumber of different bits=%d\n", hash1, hash2, countDifferentBits(hash1, hash2))

	fmt.Println()
	hash1 = sha256.Sum256([]byte("92"))
	hash2 = sha256.Sum256([]byte("30"))
	fmt.Printf("%x\n%x\nNumber of different bits=%d\n", hash1, hash2, countDifferentBits(hash1, hash2))

	fmt.Println()
	hash1 = sha256.Sum256([]byte("3"))
	hash2 = sha256.Sum256([]byte("5"))
	fmt.Printf("%x\n%x\nNumber of different bits=%d\n", hash1, hash2, countDifferentBits(hash1, hash2))

	fmt.Println()
	hash1 = sha256.Sum256([]byte("voltron"))
	hash2 = sha256.Sum256([]byte("optimus"))
	fmt.Printf("%x\n%x\nNumber of different bits=%d\n", hash1, hash2, countDifferentBits(hash1, hash2))

}

func countDifferentBits(hash1 [32]byte, hash2 [32]byte) int {
	totalDifferentBits := 0
	for i := 0; i < 32; i += 1 {
		differentBits := hash1[i] ^ hash2[i]
		totalDifferentBits += int(pc[differentBits])
	}
	return totalDifferentBits
}
