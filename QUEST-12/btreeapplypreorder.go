package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {

		f(root.Data)
		BTreeApplyInorder(root.Left, f)
		BTreeApplyInorder(root.Right, f)

	}
}
