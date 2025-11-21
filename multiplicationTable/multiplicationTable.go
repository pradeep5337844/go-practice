package main

import "fmt"

func main() {
	numToMultiply := 5
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d*%d=%d\n", numToMultiply, i, numToMultiply*i)
	}
}
