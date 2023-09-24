package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListRemoveIf(l *List, data_ref interface{}) {
	elem := l.Head
	for elem != nil {
		if elem.Data == data_ref {
			elem.Data = nil
		}
		elem = elem.Next
	}
}
