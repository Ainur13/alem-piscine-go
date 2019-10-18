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
		if rem == 1 {
			z01.PrintRune('1')
		}
		if rem == 2 {
			z01.PrintRune('2')
		}
		if rem == 3 {
			z01.PrintRune('3')
		}
		if rem == 4 {
			z01.PrintRune('4')
		}
		if rem == 5 {
			z01.PrintRune('5')
		}
		if rem == 6 {
			z01.PrintRune('6')
		}
		if rem == 7 {
			z01.PrintRune('7')
		}
		if rem == 8 {
			z01.PrintRune('8')
		}
		if rem == 9 {
			z01.PrintRune('9')
		}
		if rem == 0 {
			z01.PrintRune('0')
		}
	
	}	
}
