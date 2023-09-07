package piscine

func IsPrintable(s string) bool {
	i := 0
	for range s {
		if s[i] == '\n' {
			i++
			return false
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			i++
		} else {
			i++
		}
	}
	return true
}
