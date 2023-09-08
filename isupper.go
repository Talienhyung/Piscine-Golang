package piscine

func IsUpper(s string) bool {
	i := 0
	e := 0
	for range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			i++
			e++
		} else {
			i++
		}
	}
	return e == i
}
