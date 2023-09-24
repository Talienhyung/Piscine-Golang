package piscine

func ListMerge(l1 *List, l2 *List) {
	elem := l1.Head
	for elem != nil {
		elem = elem.Next
	}
	elem2 := l2.Head
	for elem2 != nil {
		elem.Data = elem2.Data
		elem = elem.Next
		elem2 = elem2.Next
	}
}
