package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	isupper := false
	len := 0
	for i := range args {
		i = i
		len++
	}
	if len > 2 && args[1] == "--upper" {
		isupper = true
	}
	for i, s := range args {
		if (isupper && i > 0) || (!isupper && i > 0) {
			runes := []rune(s)
			isnum := true
			for _, r := range runes {
				if r < '0' || r > '9' {
					isnum = false
				}
			}
			if !isnum {
				z01.PrintRune(' ')
			} else {
				num := 0
				for _, r := range runes {
					num *= 10
					num += int(r - 48)
				}
				if num < 1 || num > 26 {
					z01.PrintRune(' ')
				} else {
					if isupper {
						z01.PrintRune(rune(num + 64))
					} else {
						z01.PrintRune(rune(num + 96))
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
