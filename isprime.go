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

// package piscine

// func IsPrime(nb int) bool {
// 	if nb < 0 {
// 		return false
// 	}

// 	if nb <= 3 {
// 		return nb > 1
// 	}

// 	if nb%2 == 0 || nb%3 == 0 {
// 		return false
// 	}

// 	i := 5
// 	for i*i <= nb {
// 		if nb%i == 0 || nb%(i+2) == 0 {
// 			return false
// 		}
// 		i = i + 6
// 	}

// 	return true
// }
