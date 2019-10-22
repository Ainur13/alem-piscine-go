package basicjoin

func BasicJoin(strs []string) string {
	sent := ""
	for _, s  := range strs {
		sent = sent + s
	}
	return sent
}
