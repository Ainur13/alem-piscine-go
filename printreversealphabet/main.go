package main

import "github.com/01-edu/z01"

func main() {
	for iRune := 'z'; iRune >= 'a'; iRune-- {
		z01.PrintRune(iRune)
	}
	z01.PrintRune('\n')
}
