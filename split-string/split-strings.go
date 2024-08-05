package main

import (
	"fmt"
	// Make sure to adjust the import path according to your project structure
)

func main() {
	// Test cases
	testStrings := []string{"abc", "abcdef", "dawsd", "awsaws"}

	for _, str := range testStrings {
		result := Solution(str)
		fmt.Printf("Input: \"%s\" => Output: %v\n", str, result)
	}
}

func Solution(str string) []string {
	var result []string

	// Loop through the string with a step of 2
	for i := 0; i < len(str); i += 2 {
		if i+1 < len(str) {
			// Append the pair of characters
			result = append(result, str[i:i+2])
		} else {
			// Append the last character with an underscore
			result = append(result, string(str[i])+"_")
		}
	}

	return result
}

//method 2
/*func Solution(str string) []string {
	var res []string
	if len(str) % 2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i+=2 {
		res = append(res, str[i:i+2])
	}
	return res */