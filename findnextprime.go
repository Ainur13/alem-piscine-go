package piscine

func FindNextPrime(nb int) int {
	if nb > 1000000100 {
		return 0
	}
	if nb <= 1 {
		return 2
	}
	isprime := true
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			isprime = false
			i = nb/2 + 1
		}
	}
	if isprime != true {
		return FindNextPrime(nb + 1)
	}
	return nb
}
