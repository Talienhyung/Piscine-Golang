package piscine

func MakeRange(min, max int) []int {
	if min <= max {
		n := 0
		tab := make([]int, max-min)
		for i := min; i < max; i++ {
			tab[n] = i
		}
		return tab
	} else {
		tab := make([]int, max-min)
		return tab
	}
}
