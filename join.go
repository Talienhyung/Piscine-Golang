package piscine

func Join(strs []string, sep string) string {
	var join string
	i := 0
	for range strs {
		join += strs[i]
		join += sep
		i++
	}
	return join
}
