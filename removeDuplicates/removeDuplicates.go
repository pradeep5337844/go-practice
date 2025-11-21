package main

import "fmt"

func main() {
	arr := []int{1, 2, 1, 2, 3, 5, 4, 7, 6, 5}
	fmt.Println("the new list after removing duplicates is ", removeDuplicates(arr))
}
func removeDuplicates(arr []int) []int {
	result := []int{}
	seen := make(map[int]bool)

	for _, val := range arr {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}
