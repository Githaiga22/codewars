package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

// struct 

type PasswordConfig struct {
	Length  int
	Symbols bool
	Numbers bool
	Upper   bool
}

func GeneratePassword(config PasswordConfig) (string, error) {
	if config.Length <= 0 {
		return "", fmt.Errorf("password length must be greater than 0")
	}

	
	var charset = "abcdefghijklmnopqrstuvwxyz"
	if config.Upper {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
//define number limit and add them in charset
	if config.Numbers {
		charset += "0123456789"
	}
	//defined symbols container
	if config.Symbols {
		charset += "!@#$%^&*()"
	}
	
	if len(charset) == 0 {
		return "", fmt.Errorf("at least one character type must be d")
	}
// rand for generating random numbers
// seed  for initializing random number generator starting point sequence
//time used to randomise numbers  with he local current time.now
	rand.Seed(time.Now().UnixNano())
	password := make([]byte, config.Length)
	for i := range password {
		index := rand.Intn(len(charset))
		password[i] = charset[index]
	}

	return string(password), nil
	}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <length> [upper] [number] [symbol]")
		return
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil || length <= 0 {
		fmt.Println("Please provide a valid password length greater than 0.")
		return
	}

	Upper := false
	Numbers := false
	Symbols := false

	if len(os.Args) > 2 {
		Upper = strings.ToLower(os.Args[2]) == "upper"
	}
	if len(os.Args) > 3 {
		Numbers = strings.ToLower(os.Args[3]) == "number"
	}
	if len(os.Args) > 4 {
		Symbols = strings.ToLower(os.Args[4]) == "symbol"
	}

	config := PasswordConfig{
		Length:         length,
		Symbols: Symbols,
		Numbers: Numbers,
		Upper:   Upper,
	}

	password, err := GeneratePassword(config)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	fmt.Println("Generated Password:", password)

}