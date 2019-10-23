package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := range args {
		for _, l := range args[i] {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}
