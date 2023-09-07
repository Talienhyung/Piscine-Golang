package piscine

func AlphaCount(s string) int {
	i := 0
	e := 0
	for range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			i++
			e++
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			i++
			e++
		} else {
			i++
		}
	}
	return e
}
