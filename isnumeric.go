package piscine

func IsNumeric(s string) bool {
	i := 0
	for range s {
		if s[i] >= '0' && s[i] <= '9' {
			i++
		} else {
			return false
		}
	}
	return true
}
