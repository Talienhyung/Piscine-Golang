package piscine

func Atoi(s string) int {
	non := 0
	nega := 2
	for i := range s {
		if s[i] != '-' && s[i] != '+' {
			if s[i] < '0' || s[i] > '9' {
				return 0
			}
		}
		if s[i] >= '0' && s[i] <= '9' {
			non = non*10 + int(s[i]-'0')
		}
		if s[i] == '-' {
			if i == 0 {
				nega = -1
			} else {
				return 0
			}
		} else if s[i] == '+' {
			if i == 0 {
				nega = 1
			} else {
				return 0
			}
		}
	}
	if nega == 2 {
		nega = 1
	}
	return non * nega
}
