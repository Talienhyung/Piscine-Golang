package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListReverse(l *List) {
	ListeAlenvert := &List{}
	element := l.Head
	for element != nil {
		ListPushFront(ListeAlenvert, element.Data)
		element = element.Next
	}
	l.Head = ListeAlenvert.Head
	l.Tail = ListeAlenvert.Tail
}
