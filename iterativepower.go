package piscine

func IterativePower(nb int, power int) int {
	nnb := nb
	if nb == 0 {
		return 1
	} else if power >= 0 {
		for i := 0; i < power-1; i++ {
			nb *= nnb
		}
		return nb
	}
	return 0
}
