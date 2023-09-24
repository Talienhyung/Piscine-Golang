package piscine

func ListPushBack(l *List, data interface{}) {
	element := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = element
	} else {
		l.Tail.Next = element
	}
	l.Tail = element
}
