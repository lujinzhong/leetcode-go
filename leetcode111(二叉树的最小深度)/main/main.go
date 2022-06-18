package main

//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明：叶子节点是指没有子节点的节点。
//
//树中节点数的范围在 [0, 105] 内
//-1000 <= Node.val <= 1000

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法，要注意左右孩子其中一个不存在的时候，要取存在的节点
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	inOrder := func(root *TreeNode) int {
		leftHight := minDepth(root.Left)
		rightHight := minDepth(root.Right)
		// 几种情况
		//1.如果左孩子是nil,右孩子不是，要取右孩子
		//2.如果右孩子是 nil ，左孩子不是，要取左孩子
		// 如果都不是 nil ,根据大小判断
		// 如果都是 nil ,哪个都可以
		if root.Left == nil && root.Right != nil {
			return rightHight + 1
		}
		if root.Left != nil && root.Right == nil {
			return leftHight + 1
		}
		// 都是nil ，或者都不是nil,直接判断，无所谓了
		if leftHight < rightHight {
			return leftHight + 1
		} else {
			return rightHight + 1
		}
	}
	return inOrder(root)
}

func main() {

}
