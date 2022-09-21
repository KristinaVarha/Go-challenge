package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		n := nb + 1
		for IsPrime(n) == false {
			n++
		}
		return n
	}
}

func IsPrime2(nb int) bool {
	if nb == 1 || nb <= 0 {
		return false
	}
	prime := 1
	for i := 2; i < nb; i++ {
		if (nb % i) == 0 {
			return false
		}
	}
	if prime == 1 {
		return true
	} else {
		return false
	}
}
