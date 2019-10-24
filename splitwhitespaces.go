package piscine

func SplitWhiteSpaces(str string) []string {
	var arr []string
	word := ""
	i := 0
	for _, l := range str {
		if l == ' ' {
			arr[i] = word
			word = ""
			i++
		} else {
			word += string(l)
		}
	}
	arr[i] = word
	return arr
}
