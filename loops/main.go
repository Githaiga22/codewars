package main

import (
	"fmt"
	
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
func FifthAndSkip(str string) string {
if len(str) < 5 {
	return "Invalid Input" + "\n"
}

if len(str) == 0 {
	return "Invalid Input" + "\n"
}
res := ""
count := 0

 for i := 0; i < len(str); i++ {
if str[i] == ' ' {
continue
}
//add the char to the result
res+= string(str[i])
count++

if count == 5 {
	//skip the next char
	i++
	count = 0
	if i < len(str) {
		//add the space within 6th char
		res+= " "
	}

}

 }
 return res + "\n"
}