package atoi

func Atoi(s string) int {
	runes := []rune(s)
	res := 0
	m := 0
	l := 0
	for _, r := range runes {
		l++
		res *= 10
		if r == ' ' {
			return 0
		}
		if r != '-' && r != '+' && r != '0' && r != '1' && r != '2' && r != '3' && r != '4' && r != '5' && r != '6' && r != '7' && r != '8' && r != '9' {
			return 0
		}
		if r == '-' {
			m++
		}
		if (r == '-' || r == '+') && l > 1 {
			return 0
		}
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
	if m == 1 {
		res *= -1
	}
	return res
}
