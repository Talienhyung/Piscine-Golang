package piscine

func IsPrime(nb int) bool {
	if nb > 1 {
		if nb == 2 {
			return true
		}
		for i := 2; i < 100000; i++ {
			if nb%i == 0 && nb != i {
				return false
			}
		}
		return true
	}
	return false
}
