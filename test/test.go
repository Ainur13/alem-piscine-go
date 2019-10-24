package main

import "fmt"

func SplitWhiteSpaces(str string) []string {
	lens := 0
	for i := range str {
		i = i
		lens++
	}
	len := 0
	for i, l := range str {
		if l == ' ' && i > 0 && str[i-1] != ' ' && i != lens-1 {
			len++
		}
	}
	arr := make([]string, len+1)
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

func main() {
	str := " Hello  how are you? "
	fmt.Println(SplitWhiteSpaces(str))
	fmt.Println(len(SplitWhiteSpaces(str)))
}
