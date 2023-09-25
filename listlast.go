package piscine

func ListLast(l *List) interface{} {
	element := l.Tail
	for element != nil {
		element = element.Next
		if element == nil {
			return l.Tail.Data
		}
	}
	return nil
}
