package main

import (
	piscine ".."
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for _, w := range table {
		runes := []rune(w)
		for _, r := range runes {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func main() {
	str := "Hello how are you?"
	table := piscine.SplitWhiteSpaces(str)
	PrintWordsTables(table)
}
