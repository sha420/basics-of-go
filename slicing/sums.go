package main

import "fmt"

func sum(num ...int) int {
	res := 0
	for i := 0; i < len(num); i++ {
		res = res + i
	}
	return res
}
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
