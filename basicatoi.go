package piscine

func BasicAtoi(s string) int {
	runes := []rune(s)
	res := 0
	for _, r := range runes {
		res *= 10
		if r == '0' {
			res += 0
		}
		if r == '1' {
			res += 1
		}
		if r == '2' {
			res += 2
		}
		if r == '3' {
			res += 3
		}
		if r == '4' {
			res += 4
		}
		if r == '5' {
			res += 5
		}
		if r == '6' {
			res += 6
		}
		if r == '7' {
			res += 7
		}
		if r == '8' {
			res += 8
		}
		if r == '9' {
			res += 9
		}
	}
	return res
}
