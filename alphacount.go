package piscine

func AlphaCount(s string) int {
	i := 0
	for range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			i++
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			i++
		}
	}
	return i
}
