package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	elem2 := n2
	elem := n1
	for elem2 != nil {
		elem = SortListInsert(elem, elem2.Data)
		elem2 = elem2.Next
	}
	return n1
}
