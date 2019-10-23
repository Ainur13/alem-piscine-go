package main

import (
	"os"
	"strings"

	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	help := "--insert\n\t-i\n\t\tThis flag inserts the string into the string passed as argument.\n--order\n\t-o\n\t\tThis flag will behave like a boolean, if it is called it will order the argument"
	len := 0
	for i := range args {
		i = i
		len++
	}
	if len == 1 {
		fmt.Print(help)
	} else {
		order := false
		insert := ""
		for i, s := range args {
			if i > 0 {
				substrings := strings.Split(s, "=")
				if substrings[0] == "--insert" || substrings[0] == "-i" {
					insert = substrings[1]
				} else if s == "--order" || s == "-o" {
					order = true
				} else if s == "--help" || s == "-h" {
					fmt.Print(help)
				} else {
					if insert != "" {
						s += insert
					}
					if order {
						runes := []rune(s)
						len = 0
						for k := range runes {
							k = k
							len++
						}
						for k := 0; k < len-1; k++ {
							for j := k + 1; j < len; j++ {
								if runes[k] > runes[j] {
									t := runes[k]
									runes[k] = runes[j]
									runes[j] = t
								}
							}
						}
						s = string(runes)
					}
					fmt.Print(s)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
