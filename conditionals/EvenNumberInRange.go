package main

import "fmt"

func main() {
	for n := range 126 {
		if n%2 == 0 {
			fmt.Println(n) 
		}
	}
}
