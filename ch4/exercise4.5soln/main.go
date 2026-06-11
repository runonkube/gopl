package main

import "fmt"

func main() {
	names := []string{"Timi", "John", "Dave", "Dave", "James", "James", "Dave"}
	names = removeAdjacentDuplicates(names)
	fmt.Println(names)

	letters := []string{"a", "a", "a"}
	letters = removeAdjacentDuplicates(letters)
	fmt.Println(letters)
}

func removeAdjacentDuplicates(names []string) []string {
	for i := 0; i < len(names)-1; {

		if names[i] == names[i+1] {
			copy(names[i:], names[i+1:])
			names = names[:len(names)-1]
		} else {
			i += 1
		}
	}
	return names[:]
}
