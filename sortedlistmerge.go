package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	elem2 := n2
	for elem2 != nil {
		n1 = SortListInsert(n1, elem2.Data)
		elem2 = elem2.Next
	}
	return n1
}
