package main

import "fmt"

/*
func main(){
	a:= 10
	b:=20

	a,b=b,a
	fmt.Println("value of a is ", a, " value of b is ", b)
}
*/

func main() {
	a := 10
	b := 20

	a, b = swap(a, b)
	fmt.Println("value of a is ", a, " value of b is ", b)
}
func swap(a, b int) (int, int) {
	return b, a
}
