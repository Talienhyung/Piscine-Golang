package piscine

func BasicAtoi(s string) int {
	non := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			non = non*10 + int(c-'0')
		}
	}
	return non
}
