package piscine

import "fmt"

//import "github.com/01-edu/z01.PrintRune"

func EightQueens() {
	var pos [8]int
	var lines [8]bool
	var diag1 [8]bool
	var diag2 [8]bool

	x := 0
	c := 0
	for x >= 0 {
		if x == 8 || pos[x] == 8 {
			if x == 8 {
				c++
			} else {
				pos[x] = 0
			}
			x--
			lines[pos[x]] = false
			diag1[pos[x]+x] = true
			diag2[pos[x]-x+8] = false
			pos[x]++
		} else {
			if lines[pos[x]] == false && diag1[pos[x]+x] == false && diag2[pos[x]-x+8] == false {
				lines[pos[x]] = true
				diag1[pos[x]+x] = true
				diag2[pos[x]-x+8] = true
				x++
			} else {
				pos[x]++
			}
		}
	}
	fmt.Println(c)
}

func main() {
	EightQueens()
}
