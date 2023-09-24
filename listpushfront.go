package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	element := &NodeL{Data: data}
	if l.Head == nil {
		l.Tail = element
	} else {
		l.Head.Next = element
	}
	l.Head = element
}
