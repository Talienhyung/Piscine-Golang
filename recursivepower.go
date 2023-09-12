package piscine

func RecPower(nb int, power int) int {
	if nb == 0 && power == 0 {
		return 1
	} else if power == 0 {
		return 1
	} else if power > 0 {
		return nb * RecPower(nb, power-1)
	}
	return 0
}
