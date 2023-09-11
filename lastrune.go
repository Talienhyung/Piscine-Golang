package piscine

func LastRune(s string) rune {
	lene := len(s) - 1
	tab := []rune(s)
	return tab[lene]
}
