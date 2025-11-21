package main

import "fmt"

func main() {
	arr := []int{1, 2, 1, 2, 3, 5, 4, 7, 6, 5}
	fmt.Println("the frequency of elements is ", getFrequency(arr))
}
func getFrequency(arr []int) map[int]int {
	seen := make(map[int]int)

	for _, val := range arr {
		seen[val]++
	}
	return seen
}
