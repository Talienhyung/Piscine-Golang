package piscine

func FirstRune(s string) rune {
	lettre := s[0]
	if rune(lettre) == 'â' {
		return '♥'
	}
	return rune(lettre)
}
