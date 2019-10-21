package finnextprime

func FindNextPrime(nb int) int {
	if nb > 1000000 {
		return 0
	}
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
	} else {
		FindNextPrime(nb + 1)
	}
	return 0
}
