package main

//给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
//返回删除后的链表的头节点。

type ListNode struct {
	Val  int
	Next *ListNode
}

//执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗： 2.7 MB, 在所有 Go 提交中击败了 60.28% 的用户

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}

	cur := head.Next
	pre := head // 前驱节点
	for cur.Next != nil {
		if cur.Val == val {
			break
		}
		cur = cur.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return head
}

func main() {
}
