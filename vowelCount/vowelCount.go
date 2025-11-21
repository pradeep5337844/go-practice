package main

import "fmt"

/*
func main() {
	vowelsList := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	count := 0
	stringToCheck := "HEllo World"

	for _, str := range stringToCheck {
		for _, vowel := range vowelsList {
			if string(str) == vowel {
				count++
			}
		}
	}
	fmt.Println("the number of vowels in the string", stringToCheck, "is ", count)
}
*/

func main() {
	vowelsList := "aeiouAEIOU"
	count := 0
	stringToCheck := "HEllo World"

	for _, str := range stringToCheck {
		for _, vowel := range vowelsList {
			if str == vowel {
				count++
			}
		}
	}
	fmt.Println("the number of vowels in the string", stringToCheck, "is ", count)
}
