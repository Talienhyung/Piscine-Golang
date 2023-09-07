package piscine

func IsNumeric(s string) bool {
	len := 0
	for range s {
		len++
	}
	total := 0
	car := "0123456789"
	for i := 0; i < len; i++ {
		for j := 0; j < 26; j++ {
			if s[i] == car[j] {
				total++
			}
		}
	}
	return total == len
}
