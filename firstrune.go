package piscine

func FirstRune(s string) rune {
	lettre := s[0]
	if lettre == "♥" {
		lettre = '♥'
	}
	return rune(lettre)
}
