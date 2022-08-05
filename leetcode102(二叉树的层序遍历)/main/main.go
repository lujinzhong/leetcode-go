package main

import "container/list"

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//树中节点数目在范围 [0, 2000] 内
//-1000 <= Node.val <= 1000

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ { // 处理该层数据
			// 出队， 塞到临时数组
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			tmpArr = append(tmpArr, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, tmpArr)
		// 清空临时数组
		tmpArr = []int{}
	}
	return res
}

func main() {

}
