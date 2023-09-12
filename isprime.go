package piscine

func IsPrime(nb int) bool {
	if nb > 1 {
		if nb%2 == 0 {
			return false
		}
		for i := 2; i < 100000; i += 2 {
			for j := 2; j < 100000; j += 2 {
				if i*j == nb {
					return false
				}
			}
		}
		return true
	}
	return false
}
