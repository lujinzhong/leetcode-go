package main

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列

type ListNode struct {
	Val  int
	Next *ListNode
}

//执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗：2.4 MB, 在所有 Go 提交中击败了 71.91% 的用户

// 双指针迭代判断
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list2 == nil {
		return list1
	}
	if list1 == nil {
		return list2
	}

	var cur, list3 *ListNode
	if list1.Val <= list2.Val { // 先确定头
		list3 = list1
		cur = list3
		list1 = list1.Next
	} else {
		list3 = list2
		cur = list3
		list2 = list2.Next
	}
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	// 把剩余的链表元素直接都补上去
	if list1 == nil {
		cur.Next = list2
	}

	if list2 == nil {
		cur.Next = list1
	}
	return list3
}

// 递归法
func mergeTwoListsV2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsV2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsV2(list1, list2.Next)
		return list2
	}
}

func main() {

}
