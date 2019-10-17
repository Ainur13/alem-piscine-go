package strlen

func StrLen(str string) int {
	runes := []rune(str)
	len := 0
	for _, r := range runes {
		r = r
		len++
	}
	return len
}
