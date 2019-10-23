package piscine

func ToUpper(s string) string {
	res := ""
	runes := []rune(s)
	for _, r := range runes {
		if r >= 97 && r <= 122 {
			res = res + string(r-32)
		} else {
			res = res + string(r)
		}
	}
	return res
}
