package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

	if root == nil {
		return &TreeNode{Data:data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else if data == root.Data { 
		root.Parent = BTreeInsertData(root.Parent, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}
