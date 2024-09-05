package main

import "fmt"

func main() {
    fmt.Println(Unzipstring("2f"))
    fmt.Println(Unzipstring("2O5u2H0e"))
    fmt.Println(Unzipstring("3p6i1W"))
    fmt.Println(Unzipstring("6H8a"))
    fmt.Println(Unzipstring("2p4;7g"))
    fmt.Println(Unzipstring("2a 6p8f"))
    fmt.Println(Unzipstring("2t4dD"))
    fmt.Println(Unzipstring("82t4D"))
    fmt.Println(Unzipstring(""))
}
func Unzipstring(s string) string {
	//initialize the result to store str
	//initialize the length to be len(s)
	//initialize the first index

	result := ""
	length := len(s)
	i := 0

	for i < length {
		//check for valid num(0-9)
		if s[i] >= '0' && s[i] <= '9' {
			//update the count by converting it from byte to int
			count := int(s[i] - '0')
			//move to the next char
			i++
			//check if next char is a numeric alphabet
			if i < length && ((s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] < 'Z')) {
				//update our letter to currrnt char
				letter := s[i]
				//loop thru the count of times
				for j := 0; j < count; j++ {
					//update the result with the letter converted to string
					result += string(letter) 
				}
				//move to the next char
				i++
			}else {
				return "Invalid Input"
			}
			
		} else {
			return "Invalid Input"
		}
	
	}
	return result
}
