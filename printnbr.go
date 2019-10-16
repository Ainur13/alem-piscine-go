package printnbr

import "github.com/01-edu/z01"

func PrintReverse(n int) {
	if n > 0 {
		m := n % 10
		n = n / 10
		PrintReverse (n)
		z01.PrintRune(rune(m+48))
	}
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = n * (-1)
	}

	PrintReverse(n)

	z01.PrintRune(10)
}
