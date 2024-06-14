package main

import (
	"os"
	"github.com/01-edu/z01"
)

func printBinary(n int) {
	for i := 31; i >= 0; i-- {
		bit := (n >> i) & 1
		if bit == 1 {
			z01.PrintRune('*')
		} else {
			z01.PrintRune('0')
		}

		if i == 24 || i == 16 || i == 8 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func optionsToBits(args []string) int {
	var bitField int
	validFlags := "abcdefghijklmnopqrstuvwxyz"
	seenHFlag := false

	for _, arg := range args {
		if arg == "-h" {
			seenHFlag = true
			continue
		}
		if arg[0] == '-' {
			for _, c := range arg[1:] {
				if c == 'h' && seenHFlag {
					continue // Skip if 'h' has already been processed
				}
				index := -1
				for i, validChar := range validFlags {
					if rune(validChar) == c {
						index = i
						break
					}
				}

				if index == -1 {
					printInvalidOption()
					return 0
				}

				bitField |= 1 << (index)
			}
		}
	}

	return bitField
}

func printInvalidOption() {
	for _, r := range "Invalid Option" {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "-h" {
		z01.PrintRune('o')
		z01.PrintRune('p')
		z01.PrintRune('t')
		z01.PrintRune('i')
		z01.PrintRune('o')
		z01.PrintRune('n')
		z01.PrintRune('s')
		z01.PrintRune(':')
		for _, r := range " abcdefghijklmnopqrstuvwxyz" {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
		return
	}

	bitField := optionsToBits(args)
	printBinary(bitField)
}
