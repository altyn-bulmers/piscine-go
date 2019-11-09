package piscine

func BTreeLevelCount(root *TreeNode) int {
	
	if root == nil {
		return 0
	}
	
	left := BTreeLevelCount(root.Left)
	right := BTreeLevelCount(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1

}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}