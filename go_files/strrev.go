package strrev

func StrRev(s string) string {
	chars := []byte(s)
	l := 0
	for _, c := range chars {
		c = c
		l++
	}

	for i := 0; i < l/2; i++ {
		ch := chars[i]
		chars[i] = chars[l-1-i]
		chars[l-1-i] = ch
	}
	return string(chars)
}
