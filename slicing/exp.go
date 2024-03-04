package main

import "fmt"

func main() {
	a := map[int]int{1: 1, 2: 2, 3: 2, 4: 4, 5: 5, 6: 6, 7: 7, 8: 7}
	fmt.Println(a)
	for key, value := range a {
        fmt.Printf("Key: %d, Value: %d\n", key, value)
    }
}
