package piscine

func ConcatParams(args []string) string {
	res := ""
	len := 0
	for i := range args {
		i = i
		len++
	}
	for i, s := range args {
		res += s
		if len > i+1 {
			res += "\n"
		}
	}
	return res
}
