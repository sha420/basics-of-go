package main

import "fmt"

/*func sum(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println(sum(10, 20, 20))
}*/
func details() (string, int) {
	name := "Nandan"
	age := 20
	return name, age
}

func main() {
	name, age := details()
	fmt.Println(name, age)
}
