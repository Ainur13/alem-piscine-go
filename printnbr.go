package printnbr

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	m := n
	l := 0
	for m != 0 {
		l++
		m /= 10
	}
	runes := make([]rune, l)
	for i := l - 1; i >= 0; i-- {
		runes[i] = rune(n % 10)
	}
	for i := 0; i < l; i++ {
		z01.PrintRune(runes[i])
	}

}
