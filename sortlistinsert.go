package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	elem := l
	if elem.Data > data_ref {
		nouveauTruc := &NodeI{elem.Data, elem.Next}
		actuTruc := &NodeI{data_ref, nouveauTruc}
		elem = actuTruc
		return elem
	}
	ancienneL := l
	actuelL := l.Next
	for actuelL != nil {
		if actuelL.Data > data_ref {
			nouveauTruc := &NodeI{data_ref, actuelL}
			ancienneL.Next = nouveauTruc
			return l
		}
		actuelL = actuelL.Next
		ancienneL = ancienneL.Next
	}
	nouveauTruc := &NodeI{data_ref, nil}
	ancienneL.Next = nouveauTruc
	return l
}
