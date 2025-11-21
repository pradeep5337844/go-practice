package main
import "fmt"
func main(){
	for i:=0;i<=10;i++{
		if prime(i){
			fmt.Println("the number is prime", i)
		}
	}
}

func prime(num int) (isprime bool){
	if num < 2{
		return false
	}
	for i:=2;i<=num-1;i++{
		if num%i==0{
			return false
		}
	}
	return true
}
