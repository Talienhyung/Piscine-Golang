package piscine

func Capitalize(s string) string {
	car := make([]string, 0)
	i := 0
	var non string
	for range s {
		car = append(car, string(s[i]))
		i++
	}
	j := 0
	for range car {
		non = non + car[j]
		j++
	}
	return non
}
