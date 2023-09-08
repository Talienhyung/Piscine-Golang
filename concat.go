package piscine

func Concat(str1 string, str2 string) string {
	i := 0
	k := 0
	var non string
	car := make([]string, 0)
	for range str1 {
		if str1[i] >= 'A' && str1[i] <= 'Z' {
			car = append(car, string(str1[i]+32))
		} else {
			car = append(car, string(str1[i]))
		}
		i++
	}
	for range str2 {
		if str2[k] >= 'A' && str2[k] <= 'Z' {
			car = append(car, string(str2[k]+32))
		} else {
			car = append(car, string(str2[k]))
		}
		k++
	}
	j := 0
	for range car {
		non = non + car[j]
		j++
	}
	return non
}
