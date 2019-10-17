package swap

func Swap(a *int, b *int) {
	c := *a
	d := *b
	*a = d
	*b = c
}
