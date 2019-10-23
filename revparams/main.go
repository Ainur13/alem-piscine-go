package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	len := 0
	for i := range args {
		i = i
		len++
	}
	for i := len - 1; i > 0; i-- {
		for _, l := range args[i] {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}
