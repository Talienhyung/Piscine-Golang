package piscine

func BasicJoin(elems []string) string {
	var join string
	i := 0
	for range elems {
		join += elems[i]
		i++
	}
	return join
}
