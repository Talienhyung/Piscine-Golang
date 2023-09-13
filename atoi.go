package piscine

func Atoi(s string) int {
	non := 0
	nega := 2
	for _, c := range s {
		if c != '-' && c != '+' {
			if c < '0' || c > '9' {
				return 0
			}
		}
		if c >= '0' && c <= '9' {
			non = non*10 + int(c-'0')
		}
		if c == '-' && nega == 2 && c == rune(s[0]) {
			nega = -1
		} else if c == '+' && nega == 2 && c == rune(s[0]) {
			nega = 1
		} else if (c == '-' || c == '+') && (nega == -1 || nega == 1) && c != rune(s[0]+'0') {
			return 0
		}
	}
	if nega == 2 {
		nega = 1
	}
	return non * nega
}
