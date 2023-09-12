package piscine

func Capitalize(s string) string {
	var car []string
	i := 0
	var non string
	for range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			if i > 0 {
				if s[i-1] >= '0' && s[i-1] <= '9' {
					car = append(car, string(s[i]))
				} else if s[i-1] >= 'A' && s[i-1] <= 'Z' {
					car = append(car, string(s[i]))
				} else if s[i-1] >= 'a' && s[i-1] <= 'z' {
					car = append(car, string(s[i]))
				} else {
					car = append(car, string(s[i]-32))
				}
			} else {
				car = append(car, string(s[i]-32))
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			if i > 0 {
				if s[i-1] >= '0' && s[i-1] <= '9' {
					car = append(car, string(s[i]+32))
				} else if s[i-1] >= 'A' && s[i-1] <= 'Z' {
					car = append(car, string(s[i]+32))
				} else if s[i-1] >= 'a' && s[i-1] <= 'z' {
					car = append(car, string(s[i]+32))
				} else {
					car = append(car, string(s[i]))
				}
			} else {
				car = append(car, string(s[i]))
			}
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
