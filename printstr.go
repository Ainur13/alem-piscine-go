package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	characters := []byte(str)
	for _, ch := range characters {
		z01.PrintRune(rune(ch))
	}
}
