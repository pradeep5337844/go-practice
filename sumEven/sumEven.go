package main

import "fmt"

func main(){
	sum:= 0
	for i:=0;i<=100;i++{
		if i%2==0{
			sum=sum+i
		}
	}
	fmt.Println("sum of even numbers is ", sum)
}
