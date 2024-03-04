package main

import "fmt"

func main() {
	var a = 0
	if a%2 == 0 || a == 0 {
		fmt.Printf("It is even")
	} else if a%2 == 1 {
		fmt.Printf("It is odd")
	} else {
		fmt.Printf("Thank You")
	}
}
