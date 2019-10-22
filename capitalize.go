package piscine

func Capitalize(s string) string {
	res := ""
	runes := []rune(s)
	for i, r := range runes {
		if i == 0 && r >= 'a' && r <= 'z' {
			res = res + string(r-32)
		} else if i!= 0 && r >= 'A' && r <= 'Z' && (runes[i-1] >= 'a' && runes[i-1] <= 'z' || runes[i-1] >= 'A' && runes[i-1] <= 'Z' || runes[i-1] >= '0' && runes[i-1] <= '9') {
			res = res + string(r+32)
		} else if i!= 0 && r >= 'a' && r <= 'z' && (runes[i-1] < 'a' || runes[i-1] > 'z') && (runes[i-1] < 'A' || runes[i-1] > 'Z') && (runes[i-1] < '0' || runes[i-1] > '9') {
			res = res + string(r-32)
		} else {
			res = res + string(r)
		}
	}
	return res
}