package piscine

func Join(strs []string, sep string) string {
	var join string
	i := 0
	for range strs {
		join += strs[i]
		if i != len(strs)-1 {
			join += sep
		}
		i++
	}
	return join
}
