package piscine

func IsUpper(s string) bool {
	len := 0
	for range s {
		len++
	}
	total := 0
	car := "AZERTYUIOPQSDFGHJKLMWXCVBN"
	for i := 0; i < len; i++ {
		for j := 0; j < 26; j++ {
			if s[i] == car[j] {
				total++
			}
		}
	}
	return total == len
}
