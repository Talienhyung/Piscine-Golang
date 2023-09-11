package piscine

func IsUpper(s string) bool {
	i := 0
	for range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			i++
		} else {
			return false
		}
	}
	return true
}
