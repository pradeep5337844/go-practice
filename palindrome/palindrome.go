package main

import "fmt"

func main() {
	palindromecheckvalue := "madam"
	palincheck:=isPalindromeOrNot(palindromecheckvalue)
	if palincheck{
		fmt.Println("palindrome")
	}else{
		fmt.Println("not palindrome")
	}
}

func isPalindromeOrNot(v string) bool{
	fmt.Println("the given value is ", v)
	s:= []rune(v)

	for i,j:= 0,len(s)-1;i<j;i,j= i+1,j-1{
		if s[i]!=s[j]{
			return false
		}
	}
	return true
}
