package alphacount

func AlphaCount(str string) int {
	c := 0
	for _, l := range str {
		if l >= 'a' && l <= 'z' || l >= 'A' && l <= 'Z' {
			c++
		}
	}
	return c
}
