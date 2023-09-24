package piscine

func ListSize(l *List) int {
	compte := 0
	element := l.Head
	for element != nil {
		compte++
		element = element.Next
	}
	return compte
}
