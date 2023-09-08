package piscine

func ToLower(s string) string {
	i := 0
	var non string
	car := make([]string, 0)
	for range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			car = append(car, string(s[i]+32))
		} else {
			car = append(car, string(s[i]))
		}
		i++
	}
	j := 0
	for range car {
		non = non + car[j]
		j++
	}
	return non
}
