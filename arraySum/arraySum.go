package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, ele := range arr {
		sum = sum + ele
	}
	fmt.Println("the sum of array is ", sum)
}
