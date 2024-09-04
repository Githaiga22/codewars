package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}
 //access thru elements in the first string
	input := args[0]
 //initialize the slice of string to store values
	slice := []string{}
   //initialize the string to as empty
	str := ""
 //
	for i, char := range input {
		if char != ' ' { //if the first char is not empty
			str += string(char) // we add the char in the str
		}
		if char == ' ' && str != "" { //if char is empty and str has char
			slice = append(slice, str) // append slice with words in str
			str = "" // reset str to empty for the next word
		}
		if i == len(input)-1 && str != "" { // if we are at the last index of char and it has no space
			slice = append(slice, str) //append the last char(word)
		}
	}
	for i, value := range slice { //range through the slice 
		for _, letter := range value { // range through individual runes 
			z01.PrintRune(rune(letter)) // print individual runes
		}
		if i < len(slice)-1 { //if finished printing the first word in runes
			z01.PrintRune(' ') // print a space
		}
	}
	z01.PrintRune('\n')
}
