package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	elem := l.Head
	for elem != nil {
		if comp(elem.Data, ref) {
			return &elem.Data
		}
		elem = elem.Next
	}
	return nil
}
