package main
import "fmt"

func main(){
	arr:= []int{5,3,7,1,9}
	max:= arr[0]
	min:=arr[0]
	for _,val:= range arr{
		if val>max{
			max=val
		}
		if val<min{
			min=val
		}
	}
	fmt.Println("max",max)
	fmt.Println("min", min)
}