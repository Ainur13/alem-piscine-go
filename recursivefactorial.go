package recursivefactorial

func RecursiveFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	if nb > 1 {
		nb = nb * RecursiveFactorial(nb-1)
		return nb
	}
	return 0
}