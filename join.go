package piscine

func Join(strs []string, sep string) string {
	var join string
	for i := range strs {
		join += strs[i]
		if i != len(strs)-1 {
			join += sep
		}
	}
	return join
}
