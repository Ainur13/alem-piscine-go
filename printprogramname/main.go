package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	filename := os.Args[0]
	for _, l := range filename {
		z01.PrintRune(l)
	}
}