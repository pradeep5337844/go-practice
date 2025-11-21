package main

import "fmt"

func main() {
	n := 6
	fmt.Println("the fibonacci of ", n, " is ", fibonacci(n))
}
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
