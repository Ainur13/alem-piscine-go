package nrune

func NRune(s string, n int) rune {
	r := []rune(s)
	return r[n-1]
}
