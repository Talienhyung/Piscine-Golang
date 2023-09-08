package piscine

func FirstRune(s string) rune {
	lettre := s[0]           //on recupere la premiere lettre
	if rune(lettre) == 'â' { //comme le coeur ne marche pas on dit que si ca donne un â au lieu d'un coeur il faut bien y mettre un coeur
		return '♥'
	}
	return rune(lettre) //on return la lettre en la convert en rune (car a la base c'est une string)
}
