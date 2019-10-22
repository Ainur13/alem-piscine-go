package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)
	for _, r := range runes {
		if r < 32 {
			return false
		}
	}
	return true
}
