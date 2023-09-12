package piscine

func BasicJoin(elems []string) string {
	var join string
	for i := range elems {
		join += elems[i]
	}
	return join
}
