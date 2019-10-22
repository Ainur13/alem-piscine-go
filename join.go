package piscine

func Join(strs []string, sep string) string {
	sent := ""
	l := 0
	for i := range strs {
		i = i
		l++
	}
	for i, s := range strs {
		if i != l-1 {
			sent = sent + s + sep
		} else {
			sent = sent + s
		}
	}
	return sent
}
