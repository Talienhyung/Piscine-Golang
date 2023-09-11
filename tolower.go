package piscine

func ToLower(s string) string {
	i := 0
	var non string
	for range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			non = non + string(s[i]+32)
		} else {
			non = non + string(s[i])
		}
		i++
	}
	return non
}
