package piscine

func SplitWhiteSpaces(str string) []string {
	len := 0
	for _, l := range str {
		if l == ' ' {
			len++
		}
	}
	arr := make([]string, len)
	word := ""
	i := 0
	for _, l := range str {
		if l == ' ' {
			if word != "" {
				arr[i] = word
				i++
			}
			word = ""
		} else {
			word += string(l)
		}
		if word != "" {
			arr[i] = word
		}
	}
	return arr
}
