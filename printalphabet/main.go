package main

import "github.com/01-edu/z01"

func main() {
	for iRune := 'a'; iRune <= 'z'; iRune++ {
		z01.PrintRune(iRune)
	}
	z01.PrintRune('\n')
}
