package finnextprime

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	isprime := true
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			isprime = false
		}
	}
	if isprime == true {
		return nb
	}
	for j := nb; j <= nb+1000; j++ {
		isprime = true
		for i := 2; i <= j/2; i++ {
			if j%i == 0 {
				isprime = false
				i = j/2 + 1
			}
		}
		if isprime == true {
			return j
		}
	}
	return 0
}
