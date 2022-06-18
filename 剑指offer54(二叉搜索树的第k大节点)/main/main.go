package main

//给定一棵二叉搜索树，请找出其中第 k 大的节点的值。
//1 ≤ k ≤ 二叉搜索树元素个数

var output []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 利用二叉搜索树的特性，中序遍历后获取递增有序数组，直接取值
func kthLargest(root *TreeNode, k int) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	// 中序遍历获得递增的数组
	arr := inorderTraversal(root)
	return arr[len(arr)-k]
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

}
