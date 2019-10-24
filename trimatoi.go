package piscine

func TrimAtoi(s string) int {
	c := 0
	isneg := false
	num := 0
	for _, l := range s {
		if c == 0 && l == '-' {
			isneg = true
		}
		if l >= '0' && l <= '9' {
			num *= 10
			num += int(l - 48)
			c++
		}
	}
	if isneg {
		return -num
	} else {
		return num
	}
}
