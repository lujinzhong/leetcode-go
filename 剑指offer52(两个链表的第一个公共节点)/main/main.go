package main

//输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
//输出：Reference of the node with value = 8
//输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 暴力解法，先遍历一个链表,放到哈希表里，然后遍历第二个链表，看节点有没有在哈希表里
// 双指针，两个链表长度分别为L1+C、L2+C， C为公共部分的长度， 第一个人走了L1+C步后，回到第二个人起点走L2步；第2个人走了L2+C步后，回到第一个人起点走L1步。 当两个人走的步数都为L1+L2+C时就两个家伙就相爱了
// L1+C+（L2-C) = L1+L2
// L2+C+(L1-C) = L1+L2

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { // 有任意一个链表为空，都不可能有相交节点
		return nil
	}
	if headA == headB { // 一开始就在一起了
		return headB
	}

	l1 := headA
	l2 := headB
	for {
		if l1.Next == nil {
			l1 = headB // 自己的路走完了，开始走对方的道路
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2 = headA // 自己的路走完了，开始走对方的道路
		} else {
			l2 = l2.Next
		}
		if l1 == l2 { // 勇敢相遇
			return l1
		}
		if l1.Next == nil && l2.Next == nil { // 一起走到了各自的尽头
			return nil
		}
	}
}

func main() {
}
