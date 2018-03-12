package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	input = strings.Replace(input, "\r", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	// For Unix: test with delimiter "\n", for Windows: test with ""
	switch input {
	case "Philip":
		fmt.Println("Welcome Philip!")
	case "Chris":
		fmt.Println("Welcome Chris!")
	case "Ivo":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "Philip":
		fallthrough
	case "Ivo":
		fallthrough
	case "Chris":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip", "Ivo":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}
