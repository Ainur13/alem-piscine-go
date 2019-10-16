package printcomb

import "github.com/01-edu/z01"

func PrintComb() {
	for iRune := '0'; iRune <= '7'; iRune++ {
		for iRune2 := iRune + 1; iRune2 <= '9'; iRune2++ {
			for iRune3 := iRune2 + 1; iRune3 <= '9'; iRune3++ {
				z01.PrintRune(iRune)
				z01.PrintRune(iRune2)
				z01.PrintRune(iRune3)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune(10)
}