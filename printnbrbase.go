package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	l := 0
	for _, s := range base {
		if s == '-' || s == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		l++
	}
	if l < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	res := ""
	for nbr > 0 {
		rem := nbr % l
		res += string(base[rem])
		nbr /= l
	}
	ll := 0
	for i := range res {
		i = i
		ll++
	}
	for i := ll - 1; i >= 0; i-- {
		z01.PrintRune(rune(res[i]))
	}
}
