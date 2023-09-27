package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root.Data == elem {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	}
	return root
}
