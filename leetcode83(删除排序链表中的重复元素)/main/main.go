package main

//给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
//链表中节点数目在范围 [0, 300] 内
//-100 <= Node.val <= 100
//题目数据保证链表已经按升序 排列

type ListNode struct {
	Val  int
	Next *ListNode
}

// 一次遍历
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

// 双指针+基准值
func deleteDuplicatesV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //没有结点或者只有一个结点
		return head
	}
	cur, pre := head.Next, head
	base := pre.Val
	for cur != nil {
		if cur.Val != base {
			base = cur.Val
			cur = cur.Next
			pre = pre.Next
		} else {
			pre.Next = cur.Next
			cur = cur.Next
		}
	}
	return head
}

func main() {

}
