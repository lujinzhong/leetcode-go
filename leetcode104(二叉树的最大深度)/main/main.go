package main

//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明: 叶子节点是指没有子节点的节点。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	inOrder := func(root *TreeNode) int {
		leftHight := maxDepth(root.Left)
		rightHight := maxDepth(root.Right)
		if leftHight > rightHight {
			return leftHight + 1
		} else {
			return rightHight + 1
		}
	}
	return inOrder(root)

}
func main() {

}
