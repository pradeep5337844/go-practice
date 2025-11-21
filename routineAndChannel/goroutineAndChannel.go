package main

import "fmt"

func main() {
	names := []string{"pradeep", "samruddhi", "sameep"}
	namech := make(chan string)
	go func() {
		for ch := range namech {
			fmt.Println(ch)
		}
	}()
	for _, name := range names {
		printHello(name, namech)
	}
	close(namech)

}
func printHello(name string, namech chan string) {
	namech <- name
	fmt.Println("HELLO ", name)
}
