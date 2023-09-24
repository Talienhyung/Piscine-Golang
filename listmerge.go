package piscine

func ListMerge(l1 *List, l2 *List) {
	elem2 := l2.Head
	for elem2 != nil {
		ListPushBack(l1, elem2.Data)
		elem2 = elem2.Next
	}
}
