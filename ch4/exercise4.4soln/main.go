package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	b := rotate(a[:], 2)
	fmt.Println(b)

	b = rotate(a[:], 5)
	fmt.Println(b)
}

func rotate(s []int, positions int) []int {
	if positions <= 0 {
		return s
	}
	positions = positions % len(s)
	rotated := s[positions:]

	for i := 0; i < positions; i += 1 {
		rotated = append(rotated, s[i])
	}
	return rotated
}
