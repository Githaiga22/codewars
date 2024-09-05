package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	// if string is empty return empty string
if len(s) == 0 {
	return ""
}
//initialize the result to store the string
//initialize the previous char as the first index so as to compare repeating char
//initialize the count to 1 
result := ""
previouschar := s[0]
count := 1

//loop through the string from the second char
for i := 1; i < len(s); i++ {
	//initialize the current char so as to enable comparison
	currentchar := s[i]
	//if current char is same as previous char the count should increase
	if previouschar == currentchar {
		count++
		//else if they dont repeat, add the count and the previouschar in the result
		//remember to convert the count(int) to a string and previouschar(byte) to a string
	} else {
		result += strconv.Itoa(count) + string(previouschar)
		//reset count to 1 for allowing next number
		count = 1
		//update the previous char is now current char
		previouschar = currentchar
	}
}
//store now the remaining char which is left out at the end
result += strconv.Itoa(count) + string(previouschar)
return result

}
