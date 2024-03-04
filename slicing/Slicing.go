package main

import "fmt"

func main() {
	var a = make([]int, 3)
	fmt.Println(a)
	fmt.Println("Shakeel", len(a))
	fmt.Println("Shakeel Ahmed", len(a) == 4)
	var b = []int{1, 2, 3, 4}
	var c = append(b, 3)
	fmt.Println(c)
	var name = "vasavi"
	fmt.Println(name[1:4])
	var pname = "Nandan Maharshi"
	fmt.Println(pname[1:11])

}
