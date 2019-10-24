package piscine

func SortIntegerTable(table []int) {
	l := 0
	for i := range table {
		i = i
		l++
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if table[i] > table[j] {
				t := table[i]
				table[i] = table[j]
				table[j] = t
			}
		}
	}
}
