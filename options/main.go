package main

import (
	"fmt"
	"os"
	"strings"
)

// Function to convert a 32-bit integer to a formatted binary string
func intToBinaryString(n int) string {
	binaryStr := fmt.Sprintf("%032b", n) // Convert integer to a 32-bit binary string
	// Split binary string into 4 groups of 8 bits
	return fmt.Sprintf("%s %s %s %s", binaryStr[0:8], binaryStr[8:16], binaryStr[16:24], binaryStr[24:32])
}

// Function to process command-line arguments and output the result
func main() {
	args := os.Args[1:] // Get command-line arguments, excluding the program name
	// if len(args) == 0 || args[0] == "-h" {
	// 	// Handle no arguments or the -h flag
	// 	fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
	// 	return
	// }

	var result int // This will store the binary representation as an integer
	//validOptions := "abcdefghijklmnopqrstuvwxyz"

	for _, arg := range args {
		if arg == "-h" {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}

		// Check if the argument starts with a hyphen and has valid characters
		if !strings.HasPrefix(arg, "-") || len(arg) == 1 {
			fmt.Println("Invalid Option")
			return
		}

		// Remove the hyphen from the start
		optionString := arg[1:]

		for _, char := range optionString {
			if char < 'a' || char > 'z' {
				fmt.Println("Invalid Option")
				return
			}

			// Calculate the bit position
			position := char - 'a'
			result |= (1 << position) // Set the corresponding bit
		}
	}

	// Print the binary representation of the result
	fmt.Println(intToBinaryString(result))
}
