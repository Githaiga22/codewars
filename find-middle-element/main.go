package main

import (
	"fmt"
)

func Gimme(array [3]int) int {
	if (array[0] > array[1] && array[0] < array[2]) || (array[0] > array[2] && array[0] < array[1]) {
		return 0
	} else if (array[1] > array[0] && array[1] < array[2]) || (array[1] > array[2] && array[1] < array[0]) {
		return 1
	} else {
		return 2
	}
}

func main() {
	// Test cases
	fmt.Println(Gimme([3]int{2, 3, 1}))   // Output: 0
	fmt.Println(Gimme([3]int{5, 10, 14})) // Output: 1
	fmt.Println(Gimme([3]int{7, 3, 5}))   // Output: 2
	fmt.Println(Gimme([3]int{15, 25, 20})) // Output: 2
}
