package piscine

func ListClear(l *List) {
	cassageDeGueule := l.Head
	for cassageDeGueule != nil {
		cassageDeGueule.Data = nil
		cassageDeGueule = cassageDeGueule.Next
	}
	l.Head = nil
}
