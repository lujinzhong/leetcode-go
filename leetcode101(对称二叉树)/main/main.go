package main

//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//树中节点数目在范围 [1, 1000] 内
//-100 <= Node.val <= 100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil { // 只有一个节点
		return true
	}
	return isSymmetryNode(root.Left, root.Right)
}

func isSymmetryNode(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil || q == nil) || p.Val != q.Val { // 父亲节点不同也不行
		return false
	}

	if (p.Left == nil && q.Right != nil) || (p.Left != nil && q.Right == nil) || (p.Right == nil && q.Left != nil) || (p.Right != nil && q.Left == nil) { // 要么都是niL，要么都是非nil
		return false
	}

	if p.Left != nil && q.Left != nil { // 如果是非nil ,判断值
		if (p.Left.Val != q.Right.Val) || (p.Right.Val != q.Left.Val) {
			return false
		}
	}
	// 重新递归
	return isSymmetryNode(p.Left, q.Right) && isSymmetryNode(p.Right, q.Left)
}

func main() {

}
