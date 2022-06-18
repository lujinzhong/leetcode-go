package main

//你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//两棵树上的节点数目都在范围 [0, 100] 内
//-104 <= Node.val <= 104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
// 判断两个树是否相同，等效于递归的判断每个节点是否相同。
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) // 只有有一个节点判断不同就不是相同树
}

func main() {

}
