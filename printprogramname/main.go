package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	filename := os.Args[0]
	for _, l := range filename {
		z01.PrintRune(l)
	}
}
