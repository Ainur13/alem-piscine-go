package nrune

func NRune(s string, n int) rune {
	r := []rune(s)
	l := 0
	for i := range s {
		i = i
		l++
	}
	if l > n {
		return r[n-1]
	} else {
		return 0
	}
}
