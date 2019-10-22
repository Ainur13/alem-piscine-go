package piscine

func Concat(str1 string, str2 string) string {
	var res []rune
	runes1 := []rune(str1)
	for _, r := range runes1 {
		res = append(res, r)
	} 
	runes2 := []rune(str2)
	for _, r := range runes2 {
		res = append(res, r)
	}
	return string(res)
}
