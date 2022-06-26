package main

//给你一个链表的头节点 head ，判断链表中是否有环。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
//
//如果链表中存在环 ，则返回 true 。 否则，返回 false 。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.直接遍历删除，需要考虑头结点
func removeElements(head *ListNode, val int) *ListNode {
	// 先把头符合的先干掉
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
	}
	if head == nil { // 没有结点
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return head
}

// 增加哨兵结点作为头
func removeElementsV2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{ // 假结点
		Val:  0,
		Next: head,
	}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return dummyHead.Next
}

// 递归法
func removeElementsV3(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElementsV3(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}

func main() {

}
