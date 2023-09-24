package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	elem := l
	for elem != nil {
		elem = elem.Next
	}
	elem.Data = data_ref
	return ListSort(l)
}
