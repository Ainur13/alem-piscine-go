package printcomb2

import "github.com/01-edu/z01"

func PrintComb2() {
	iRune := '0'
	iRune2 := '0'
	iRune3 := '0'
	iRune4 := '0'
	for i := 0; i <= 9; i++ {
		for i2 := 0; i2 <= 9; i2++ {
			for i3 := 0; i3 <= 9; i3++ {
				for i4 := 0; i4 <= 9; i4++ {
					if i*10+i2 < i3*10+i4{
						z01.PrintRune(iRune)
						z01.PrintRune(iRune2)
						z01.PrintRune(' ')
						z01.PrintRune(iRune3)
						z01.PrintRune(iRune4)

						if i != 9 || i2 != 8 || i3 != 9 || i4 != 9{
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}

					if iRune4=='9'{
						iRune4='0'
					} else {
						iRune4++
					}
				}
				if iRune3=='9'{
					iRune3='0'
				} else {
					iRune3++
				}
			}
			if iRune2=='9'{
				iRune2='0'
			} else {
				iRune2++
			}
		}
		if iRune=='9'{
			iRune='0'
		} else {
			iRune++
		}
	}
	z01.PrintRune(10)
}
