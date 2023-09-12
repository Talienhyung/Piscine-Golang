package piscine

func IterativeFactorial(nb int) int {
	r := 1
	if nb <= 2147483647 {
		if nb < 0 {
			return 0
		} else if nb == 0 {
			return 1
		} else {
			for i := 1; i <= nb; i++ {
				r *= i
			}
			return r
		}
	}
	return 0
}
