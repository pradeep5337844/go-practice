package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Int: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Bool: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
func main() {
	do(21)
	do("hello")
	do(true)
	do(20.10)
}
