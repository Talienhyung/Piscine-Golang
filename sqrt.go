package piscine

func Sqrt(nb int) int {
	for i := 1; i <= 10; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
