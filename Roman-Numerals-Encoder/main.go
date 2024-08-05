package main

import (
	"fmt"
	
)

func main() {
	// Test cases
	fmt.Println(Solution(1))    // Output: "I"
	fmt.Println(Solution(4))    // Output: "IV"
	fmt.Println(Solution(9))    // Output: "IX"
	fmt.Println(Solution(58))   // Output: "LVIII"
	fmt.Println(Solution(1990)) // Output: "MCMXC"
	fmt.Println(Solution(2024)) // Output: "MMXXIV"
	fmt.Println(Solution(3999)) // Output: "MMMCMXCIX"
}

/*
func Solution(number int) string {
	// Map of integer values to corresponding Roman numerals
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	
	// Result string to store the Roman numeral
	var result string
	
	// Iterate over the values and build the Roman numeral
	for i := 0; i < len(values); i++ {
		for number >= values[i] {
			number -= values[i]
			result += romans[i]
		}
	}
	return result
}
*/


func Solution(number int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	
	return thousands[number/1000] + hundreds[(number%1000)/100] + tens[(number%100)/10] + ones[number%10]
}
