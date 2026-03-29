package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Convert(value string, base string) {
	switch strings.ToLower(base) {
	case "hex":
		{
			val, err := strconv.ParseInt(value, 16, 64)
			if err != nil {
				fmt.Println("Error: invalid hexadecimal number! Try again!")
				return
			}
			fmt.Println("Decimal: ", val)
		}
	case "bin":
		{
			val, err := strconv.ParseInt(value, 2, 64)
			if err != nil {
				fmt.Println("Error: invalid binary number! Try again!")
				return
			}
			fmt.Println("Decimal: ", val)
		}
	case "dec":
		{
			val, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error: invalid decimal number! Try again!")
				return
			}
			fmt.Println("Binary: ", strconv.FormatInt(int64(val), 2))
			fmt.Println("Hex: ", strings.ToUpper(strconv.FormatInt(int64(val), 16)))
		}
	default:
		fmt.Println("Error: invalid base number (Use hex, bin or dec)")
	}
}

func main() {
	fmt.Println("Welcome to CLI base converter")
	Convert("1E", "hex")
	Convert("10", "bin")
	Convert("255", "dec")
}
