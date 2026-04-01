package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to CLI base converter")
	fmt.Println("Usage: convert <value> <base>")
	fmt.Println("Example: convert 1E hex")
	fmt.Println("Type 'quit' to exit program")
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "quit" {
			fmt.Println("GoodBye!")
			break
		}
		parts := strings.Fields(input)

		if len(parts) != 3 {
			fmt.Println("Unknown command. Use convert <value> <base> ")
			continue
		}

		command := strings.ToLower(parts[0])
		value := parts[1]
		base := parts[2]

		if command != "convert" {
			fmt.Println("Unknown command. Use 'convert'")
			continue
		}
		Convert(value, base)
	}

}
