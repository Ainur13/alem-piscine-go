package basicatoi

func BasicAtoi(s string) int {
	runes := []rune(s)
	l := 0
	for i := range runes {
		i = i
		l++
	}
	res := 0
	for _, r := range runes {
		st := 1
		for i := 1; i < l; i++ {
			st *= 10
		}
		res = res + int(r-48)*st
		l--
	}
	return res
}
