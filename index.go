package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	tab := []byte(s)
	atrouver := []byte(toFind)
	for i := range tab {
		if tab[i] == atrouver[0] {
			found := true
			for j := 1; j < len(atrouver); j++ {
				if i+j >= len(tab) || tab[i+j] != atrouver[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}
