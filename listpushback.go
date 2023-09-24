package piscine

func ListPushBack(l *List, data interface{}) {
	element := &NodeL{Data: data}
	if l.Head == nil {
		l.Tail = element
	} else {
		l.Head.Next = element
	}
	l.Head = element
}
