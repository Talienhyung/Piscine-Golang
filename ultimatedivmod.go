package piscine

func UltimateDivMod(a *int, b *int) {
	n := *a / *b
	n2 := *a % *b
	*a = n
	*b = n2
}
