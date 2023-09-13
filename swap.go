package piscine

func Swap(a *int, b *int) {
	y := *a
	z := *b
	*a = z
	*b = y
}
