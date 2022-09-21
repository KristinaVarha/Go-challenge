package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lastSlashIndex := 0
	programName := os.Args[0]
	for index, char := range programName {
		if char == '/' {
			lastSlashIndex = index
		}
	}

	for _, char := range programName[lastSlashIndex+1:] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
