package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the valid options as a string for simplicity
	validOptions := "abcdefghijklmnopqrstuvwxyz"

	// Initialize the options bitmask as a 32-bit integer
	var options uint32 = 0

	// Check if no arguments or the -h flag is present
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	for _, arg := range os.Args[1:] {
		if arg == "-h" {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
		if len(arg) > 1 && arg[0] == '-' {
			for i := 1; i < len(arg); i++ {
				char := arg[i]
				if char >= 'a' && char <= 'z' {
					options |= 1 << (char - 'a')
				} else {
					fmt.Println("Invalid Option")
					return
				}
			}
		} else {
			fmt.Println("Invalid Option")
			return
		}
	}

	// Print the options bitmask as groups of bytes
	fmt.Printf("%08b %08b %08b %08b\n",
		options>>24&0xFF,
		options>>16&0xFF,
		options>>8&0xFF,
		options&0xFF)
}

$ go run . | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -abc -ijk | cat -e
00000000 00000000 00000111 00000111$
$ go run . -z | cat -e
00000010 00000000 00000000 00000000$
$ go run . -abc -hijk | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -h | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -zh | cat -e
00000010 00000000 00000000 10000000$
$ go run . -z -h | cat -e
options: abcdefghijklmnopqrstuvwxyz$
