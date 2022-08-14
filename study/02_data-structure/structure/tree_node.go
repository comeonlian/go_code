package structure

import "fmt"

// TreeNode 二叉树结构定义
type TreeNode struct {
	Left, Right *TreeNode
	Data        string
}

// PreOrder 先序遍历
func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}

	// 打印根节点
	fmt.Printf("%v ", node.Data)
	// 打印左子树
	PreOrder(node.Left)
	// 打印右子树
	PreOrder(node.Right)
}

// MidOrder 中序遍历
func MidOrder(node *TreeNode) {
	if node == nil {
		return
	}

	// 打印左子树
	MidOrder(node.Left)
	// 打印根节点
	fmt.Printf("%v ", node.Data)
	// 打印右子树
	MidOrder(node.Right)
}

// PostOrder 后序遍历
func PostOrder(node *TreeNode) {
	if node == nil {
		return
	}

	// 打印左子树
	PostOrder(node.Left)
	// 打印右子树
	PostOrder(node.Right)
	// 打印根节点
	fmt.Printf("%v ", node.Data)
}
