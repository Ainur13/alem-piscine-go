package strlen

func StrLen(str string) int {
	characters := []byte(str)
	len := 0
	for _,ch := range characters {
		ch = ch
		len++
	}
	return len
}
