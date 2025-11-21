package main

import "fmt"

func main() {
	fmt.Println("year 2024 is it a leap year", isLeapYear(2024))
}

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
