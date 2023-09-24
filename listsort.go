package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	elem := l
	for elem != nil {
		elemNext := elem.Next
		for elemNext != nil {
			if elem.Data > elemNext.Data {
				elem.Data, elemNext.Data = elemNext.Data, elem.Data
			}
			elemNext = elemNext.Next
		}
		elem = elem.Next
	}
	return l
}
