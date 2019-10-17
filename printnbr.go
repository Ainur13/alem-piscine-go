package printnbr

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = n * (-1)
	}

	nn := n
	l := 0

	for nn != 0 {
		l++
		nn = nn / 10
	}

	nn = n
	rev := 0
	for nn != 0 {
		rem := nn % 10
		nn = nn / 10
		for i := 1; i < l; i++ {
			rem = rem * 10
		}
		rev = rev + rem
		l--
	}

	for rev != 0 {
		rem := rev % 10
		rev = rev / 10
		z01.PrintRune(rune(48 + rem))
	}

}
