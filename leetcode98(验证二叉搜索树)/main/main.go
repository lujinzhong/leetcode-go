package main

import "math"

//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
//有效 二叉搜索树定义如下：
//
//节点的左子树只包含 小于 当前节点的数。
//节点的右子树只包含 大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。

var output []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. 先中序遍历，如果是二叉搜索树的话，遍历出来的结果应该是一个递增的数组
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	sortArr := inorderTraversal(root)
	// 判断是不是升序的数组
	for k, v := range sortArr {
		if k != 0 {
			if v <= sortArr[k-1] { // 如果是有相等的，也不是二叉搜索树
				return false
			}
		}
	}
	return true
}

// 2. 利用递归的特性，二叉搜索树，需要满足每个子树的左孩子的值都小于父节点，父节点小于右孩子，相当于说，
//指定根节点，这个根节点的左子树的所有值必然都是比根节点小的（也就是左子树的最大值都比根值小），右子树必然都是比根节点大的（也就是说右子树的最小值都比根大）
// 每次更新对应的上下限
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var subBst func(node *TreeNode, min, max int) bool
	subBst = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return subBst(node.Left, min, node.Val) && subBst(node.Right, node.Val, max)
	}
	return subBst(root, math.MinInt64, math.MaxInt64)
}

// 中序遍历
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

}
