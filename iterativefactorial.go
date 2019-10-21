package iterativefactorial

func IterativeFactorial(nb int) int {
	res := 1
	for i := nb; i >= 1; i-- {
		res = res * i
	}
	return res
}
