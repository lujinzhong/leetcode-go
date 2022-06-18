package main

import "fmt"

//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//提示：
//
//树中节点数目在范围 [0, 100] 内
//-100 <= Node.val <= 100

var output []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法，中序遍历（左根右）
func inorderTraversal(root *TreeNode) []int {
	output = []int{}
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 左根右
		inOrder(root.Left)
		output = append(output, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return output
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(inorderTraversal(root))
}
