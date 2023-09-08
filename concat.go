package piscine

func Concat(str1 string, str2 string) string {
	i := 0
	k := 0
	var non string
	car := make([]string, 0)
	for range str1 {
		car = append(car, string(str1[i]))
		i++
	}
	for range str2 {
		car = append(car, string(str2[k]))
		k++
	}
	j := 0
	for range car {
		non = non + car[j]
		j++
	}
	return non
}
