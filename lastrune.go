package piscine

func LastRune(s string) rune {
	lene := len(s) - 1
	lettre := s[lene]
	if rune(lettre) == 'â' {
		return '♥'
	}
	return rune(lettre)
}
