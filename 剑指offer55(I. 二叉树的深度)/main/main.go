package main

//输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func maxDepth(root *TreeNode) int {
	// 递归子问题，给定一个树，求左右子树的最大深度+1
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(p, q int) int {
	if p > q {
		return p
	}
	return q
}
func main() {

}
