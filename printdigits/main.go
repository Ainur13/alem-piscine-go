package main

import "github.com/01-edu/z01"

func main() {
	for iRune := '0'; iRune <= '9'; iRune++ {
		z01.PrintRune(iRune)
	}
	z01.PrintRune('\n')
}