package printnbrinorder

import "fmt"

func PrintNbrInOrder(n int) {
	var runes []rune
	if n == 0 || n < 0 {
		fmt.Print(0)
	} else {
		l := 0
		for n > 0 {
			rem := n % 10
			runes = append(runes, rune(rem))
			n = n / 10
			l++
		}
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				if runes[i] > runes[j] {
					t := runes[i]
					runes[i] = runes[j]
					runes[j] = t
				}
			}
		}
		for _, l := range runes {
			fmt.Print(l)
		}
	}
}
