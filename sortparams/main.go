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
	for i := 1; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if args[i] > args[j] {
				t := args[i]
				args[i] = args[j]
				args[j] = t
			}
		}
	}
	for i := range args {
		if i != 0 {
			for _, l := range args[i] {
				z01.PrintRune(l)
			}
			z01.PrintRune('\n')
		}
	}
}
