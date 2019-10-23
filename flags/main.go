package main

import (
	"os"

	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	help := "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument."
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
				runes := []rune(s)
				s1 := ""
				s2 := ""
				isEquality := false
				for _, r := range runes {
					if r == '=' {
						isEquality = true
					} else if !isEquality {
						s1 += string(r)
					} else if isEquality {
						s2 += string(r)
					}
				}
				if s1 == "--insert" || s1 == "-i" {
					insert = s2
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
