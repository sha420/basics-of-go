package main

import "fmt"

func main() {
	array := [3][3]int{{1, 2, 3}, {4, 4, 5}, {72, 2, 2}}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Printf("%d\n", array[i][j])
		}
	}
}
