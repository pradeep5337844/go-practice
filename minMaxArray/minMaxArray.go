package main

import "fmt"

func main() {
	arr := []int{5, 1, 2, 7, 4, 3, 9, 8, 6}
	min := arr[0]
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println("max", max)
	fmt.Println("min", min)
}
