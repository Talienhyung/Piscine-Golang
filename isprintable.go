package piscine

func IsPrintable(s string) bool {
	i := 0
	for range s {
		if s[i] == '\n' || s[i] == '\b' || s[i] == '\r' || s[i] == '\v' || s[i] == '\f' || s[i] == '\a' || s[i] == '\t' {
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
