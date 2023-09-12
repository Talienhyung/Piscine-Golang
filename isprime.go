package piscine

func IsPrime(nb int) bool {
	if nb > 1 {
		for i := 2; i < 100; i++ {
			for j := 2; j < 100; j++ {
				if i*j == nb {
					return false
				}
			}
		}
	}
	return true
}
