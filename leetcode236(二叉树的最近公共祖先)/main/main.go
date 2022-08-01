package main

//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
//树中节点数目在范围 [2, 105] 内。
//-109 <= Node.val <= 109
//所有 Node.val 互不相同 。
//p != q
//p 和 q 均存在于给定的二叉树中。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 满足公共节点的要求，必然是 p,q 分别在该公共节点的左右子树中
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q { // 如果树为空，直接返回 nil, 如果当前节点能找到p or q , 则返回 p 或者 q
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

func main() {

}
