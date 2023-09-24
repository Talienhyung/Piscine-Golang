package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

func ListAt(l *NodeL, pos int) *NodeL {
	compter := 0
	for l != nil {
		if compter == pos {
			return l
		}
		l = l.Next
		compter++
	}
	return nil
}
