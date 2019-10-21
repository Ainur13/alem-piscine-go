package iterativefactorial

func IterativeFactorial(nb int) int {
	if nb > 20 {
		return 0
	}
	res := 1
	for i := nb; i >= 1; i-- {
		res = res * i
	}
	return res
}
