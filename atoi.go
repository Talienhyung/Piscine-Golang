package piscine

func Atoi(s string) int {
	non := 0
	nega := 1
	for _, c := range s {
		if c >= '0' && c <= '9' {
			non = non*10 + int(c-'0')
		}
		if c == '-' && non == 0 {
			nega = -1
		}
	}
	return non * nega
}
