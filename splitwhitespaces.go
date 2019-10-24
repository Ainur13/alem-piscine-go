package piscine

func SplitWhiteSpaces(str string) []string {
	var arr []string
	word := ""
	for _, l := range str {
		if l == ' ' {
			arr = append(arr, word)
			word = ""
		} else {
			word += string(l)
		}
	}
	arr = append(arr, word)
	return arr
}
