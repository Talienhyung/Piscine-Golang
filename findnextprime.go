package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	trouver := true
	for trouver {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
	return 0
}
