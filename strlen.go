package piscine

func StrLen(s string) int {
	longueur := len(s)
	longueur2 := 0
	table := "AZERTYUIOPQSDFGHJKLMWXCVBNazertyuiopqsdfghjklmwxcvbn!"
	for i := 0; i <= longueur; i++ {
		for j := 0; j < 53; j++ {
			if s[i] == table[j] {
				longueur2++
			}
		}
	}
	return longueur2
}
