package main

import (
	"fmt"
)

// Reverse reverses a string rune-wise.
func Reverse(s string) string {
	// Convert the string to a slice of runes to handle Unicode characters correctly.
	runes := []rune(s)

	// Iterate from both ends of the slice, swapping elements.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the reversed rune slice back to a string.
	return string(runes)
}

func main() {
	originalString := "Pradeep"
	reversedString := Reverse(originalString)
	fmt.Printf("Original: %s\n", originalString)
	fmt.Printf("Reversed: %s\n", reversedString)
}
