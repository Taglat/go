https://github.com/01-edu/public/blob/master/subjects/options/README.md

package main

import (
	"fmt"
	"os"
)

func main() {
	// Valid options
	validOptions := "abcdefghijklmnopqrstuvwxyz"
	helpMessage := "options: " + validOptions

	// If no arguments or "-h" flag is provided, print the help message
	if len(os.Args) < 2 || hasHFlag(os.Args[1:]) {
		fmt.Println(helpMessage)
		return
	}

	// Initialize the options bitmask
	var options int32 = 0

	// Process each argument
	for _, arg := range os.Args[1:] {
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		for i := 1; i < len(arg); i++ {
			char := arg[i]
			if !isValidOption(char) {
				fmt.Println("Invalid Option")
				return
			}

			// Set the corresponding bit in the options bitmask
			options |= 1 << (char - 'a')
		}
	}

	// Print the options bitmask as groups of bytes
	fmt.Printf("%08b %08b %08b %08b\n",
		(options>>24)&0xFF,
		(options>>16)&0xFF,
		(options>>8)&0xFF,
		options&0xFF,
	)
}

// Helper function to check if the "-h" flag exists in the arguments
func hasHFlag(args []string) bool {
	for _, arg := range args {
		if len(arg) > 1 && arg[0] == '-' {
			for i := 1; i < len(arg); i++ {
				if arg[i] == 'h' {
					return true
				}
			}
		}
	}
	return false
}

// Helper function to check if a character is a valid option
func isValidOption(char byte) bool {
	return char >= 'a' && char <= 'z'
}