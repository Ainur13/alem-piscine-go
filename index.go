package index

func Index(s string, toFind string) int {
	l1 := 0
	l2 := 0
	for i := range s {
		i = i
		l1++
	}
	for i := range toFind {
		i = i
		l2++
	}
	if l1 == 0 || l2 == 0 {
		return -1
	} 
	for i := 0; i < l1-l2; i++ {
		if s[i] == toFind[0] {
			eq := true
			for j := 1; j < l2; j++ {
				if i+j < l1 && s[i+j] != toFind[j] {
					eq = false
				}
			}
			if eq {
				return i
			}
		}
	}
	return -1
}
