package piscine

func Compare(a, b string) int {
	l1 := 0
	l2 := 0
	for i := range a {
		i = i
		l1++
	}
	for i := range b {
		i = i
		l2++
	}
	if l1 == 0 && l2 == 0 {
		return 0
	} else if l1 == 0 {
		return -1
	} else if l2 == 0 {
		return 1
	} else {
		min := 0
		if l1 >= l2 {
			min = l2
		} else if l2 > l1 {
			min = l1
		}
		for i := 0; i < min; i++ {
			if a[i] > b[i] {
				return 1
			} else if a[i] < b[i] {
				return -1
			}
		}
		if l1 > l2 {
			return 1
		} else if l2 > l1 {
			return -1
		}
	}
	return 0
}
